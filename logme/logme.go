package logme

import (
	"bytes"
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/levels"
	"log"
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
	wg           sync.WaitGroup
	buf          bytes.Buffer
	fields       []Field
}

type Settings struct {
	timestampFormat string
	mu              sync.RWMutex
}

func NewLogMe() *LogMe {
	settings := &Settings{
		timestampFormat: defaultTimestampLayout,
	}
	levelSetter := &LevelSetter{
		level: levels.All,
	}

	return &LogMe{
		outputFabric: NewOutputFabric(),
		settings:     settings,
		publishChan:  make(chan entity.Message),
		levelSetter:  levelSetter,
		active:       false,
	}
}

// Calls by level

// get timestamp as early as possible

func (lm *LogMe) Info(args ...interface{}) {
	ts := getTimestamp()
	lm.pipeline(args, ts, levels.Info)
}

func (lm *LogMe) Warning(args ...interface{}) {
	ts := getTimestamp()
	lm.pipeline(args, ts, levels.Warning)
}

func (lm *LogMe) Debug(args ...interface{}) {
	ts := getTimestamp()
	runtimeMetrics := GetRuntimeMetrics()
	text := fmt.Sprintf( // ToDo manually add/remove parameters, no hardcoding
		" | cpu %d | calls %d | gorutines %d | alloc %d | total alloc %d",
		runtimeMetrics.NumCPU, runtimeMetrics.CgoCalls, runtimeMetrics.NumGoroutine,
		runtimeMetrics.Alloc, runtimeMetrics.TotalAlloc)
	args = append(args, text)
	lm.pipeline(args, ts, levels.Debug)
	//params := dto.MessageDTO{
	//	Level: levels.Debug,
	//	Text: message + fmt.Sprintf( // ToDo manually add/remove parameters, no hardcoding
	//		" | cpu %d | calls %d | goroutines %d | alloc %d | total alloc %d",
	//		runtimeMetrics.NumCPU, runtimeMetrics.CgoCalls, runtimeMetrics.NumGoroutine,
	//		runtimeMetrics.Alloc, runtimeMetrics.TotalAlloc),
	//	Timestamp: timeStamp,
	//}
	//lm.messageProducer.NewMessage(&params)
}

func (lm *LogMe) Critical(args ...interface{}) {
	ts := getTimestamp()
	lm.pipeline(args, ts, levels.Critical)
}

func (lm *LogMe) AddOutput() *OutputFabric {
	return lm.outputFabric
}

func (lm *LogMe) sendToOutputs(message *entity.Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recovered panic from writing to %s \n", err)
		}
	}()

	outs := lm.outputFabric.GetOutputs()

	lm.wg.Add(len(outs))

	for _, o := range outs {
		go func() {
			defer lm.wg.Done()
			if err := o.Write(message); err != nil {
				panic(err)
			}
		}()
	}

	lm.wg.Wait()
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

func (lm *LogMe) pipeline(args []interface{}, timestamp time.Time, level levels.LogLevel) {
	if !lm.validate(args, level) {
		return
	}

	if len(lm.fields) != 0 {
		for _, h := range lm.fields {
			hookString := fmt.Sprintf(" || %s=%s", h.title, h.value)
			args = append(args, hookString)
		}
	}

	_, err := fmt.Fprint(&lm.buf, args...)

	if err != nil {
		return
	}

	defer lm.buf.Reset()

	timestampFormatted := timestamp.Format(lm.settings.timestampFormat)

	lm.sendToOutputs(&entity.Message{
		Text:      lm.buf,
		Level:     level,
		Timestamp: timestampFormatted,
	})
}

func (lm *LogMe) validate(message []interface{}, level levels.LogLevel) bool {
	if len(message) == 0 || lm.levelSetter.level < level {
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

func (lm *LogMe) AddHook(title string, value interface{}) {
	lm.fields = append(lm.fields, Field{title: title, value: value})
}

type Field struct {
	title string
	value interface{}
}
