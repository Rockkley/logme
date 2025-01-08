package logme

import (
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/levels"
	"strings"
	"sync"
	"time"
)

const (
	defaultTimestampLayout = time.DateTime
)

type LogMe struct {
	outputFabric *OutputFabric
	levelSetter  *LevelSetter
	settings     *Settings
	publishChan  chan entity.Message
	active       bool
}

type Settings struct {
	timestampFormat string
	mu              sync.RWMutex
}

func NewLogMe() *LogMe {
	settings := Settings{
		timestampFormat: defaultTimestampLayout,
	}
	levelSetter := LevelSetter{
		level: levels.All,
	}

	return &LogMe{
		outputFabric: NewOutputFabric(),
		settings:     &settings,
		publishChan:  make(chan entity.Message),
		levelSetter:  &levelSetter,
		active:       false,
	}
}

// Calls by level

func (lm *LogMe) Info(message string) {
	ts := getTimestamp() // get timestamp as early as possible
	lm.pipeline(message, ts, levels.Info)
}

func (lm *LogMe) Warning(message string) {
	ts := getTimestamp() // get timestamp as early as possible
	lm.pipeline(message, ts, levels.Warning)
}

func (lm *LogMe) Debug(message string) {
	//ts := getTimestamp() // get timestamp as early as possible
	//runtimeMetrics := GetRuntimeMetrics()
	//lm.newMessage(message, ts, levels.Info)
	//params := dto.MessageDTO{
	//	Level: levels.Debug,
	//	Text: message + fmt.Sprintf( // ToDo manually add/remove parameters, no hardcoding
	//		" | cpu %d | calls %d | gorutines %d | alloc %d | total alloc %d",
	//		runtimeMetrics.NumCPU, runtimeMetrics.CgoCalls, runtimeMetrics.NumGoroutine,
	//		runtimeMetrics.Alloc, runtimeMetrics.TotalAlloc),
	//	Timestamp: timeStamp,
	//}
	//lm.messageProducer.NewMessage(&params)
}

func (lm *LogMe) Critical(message string) {
	ts := getTimestamp() // get timestamp as early as possible
	lm.pipeline(message, ts, levels.Critical)
}

func (lm *LogMe) AddOutput() *OutputFabric {
	return lm.outputFabric
}

func (lm *LogMe) sendToOutputs(message *entity.Message) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovered panic from writing to %s \n", err)
		}
	}()

	var wg sync.WaitGroup

	outs := lm.outputFabric.GetOutputs()

	wg.Add(len(outs))

	for _, o := range outs {
		go func() {
			defer wg.Done()
			if err := o.Write(message); err != nil {
				fmt.Println(err)
				panic(err)
			}
		}()
	}

	wg.Wait()
}

func (lm *LogMe) SetLevel() *LevelSetter {
	return lm.levelSetter
}

func (lm *LogMe) SetTimestampLayout(timestampLayout string) {
	lm.settings.timestampFormat = timestampLayout
}

func getTimestamp() time.Time {
	return time.Now()
}

func (lm *LogMe) pipeline(message string, timestamp time.Time, level levels.LogLevel) {

	if !lm.validate(message, level) {
		return
	}

	timestampFormatted := timestamp.Format(lm.settings.timestampFormat)

	lm.sendToOutputs(&entity.Message{
		Text:      message,
		Level:     level,
		Timestamp: timestampFormatted,
	})
}

func (lm *LogMe) validate(message string, level levels.LogLevel) bool {
	if strings.TrimSpace(message) == "" || lm.levelSetter.level < level {
		return false
	}

	return true
}

func (lm *LogMe) observePublishChan() {
	lm.active = true
	for msg := range lm.publishChan {
		lm.sendToOutputs(&msg)
	}
}
