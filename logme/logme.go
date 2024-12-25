package logme

import (
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	dto2 "github.com/rockkley/logme/logme/entity/dto"
	"github.com/rockkley/logme/logme/entity/levels"
	"runtime"
	"time"
)

type LogMe struct {
	messageProducer *entity.MessageProducer
}

func NewLogMe() *LogMe {
	return &LogMe{
		messageProducer: entity.NewMessageProducer(),
	}
}

// Calls by level

func (lm *LogMe) Info(message string) {
	ts := getTimestamp()
	dto := dto2.MessageDTO{
		Level:     levels.Info,
		Text:      message,
		Timestamp: ts,
	}
	lm.messageProducer.NewMessage(&dto)
}

func (lm *LogMe) Warning(message string) {
	ts := getTimestamp()
	dto := dto2.MessageDTO{
		Level:     levels.Warning,
		Text:      message,
		Timestamp: ts,
	}
	lm.messageProducer.NewMessage(&dto)
}

func (lm *LogMe) Debug(message string) {
	ts := getTimestamp()
	metrics := GetRuntimeMetrics()
	dto := dto2.MessageDTO{
		Level: levels.Debug,
		Text: message + fmt.Sprintf(
			" | cpu %d | calls %d | gorutines %d | alloc %d | total alloc %d",
			metrics.NumCPU, metrics.CgoCalls, metrics.NumGoroutine, metrics.Alloc, metrics.TotalAlloc),
		Timestamp: ts,
	}
	lm.messageProducer.NewMessage(&dto)
}

func (lm *LogMe) Critical(message string) {
	ts := getTimestamp()
	dto := dto2.MessageDTO{
		Level:     levels.Critical,
		Text:      message,
		Timestamp: ts,
	}
	lm.messageProducer.NewMessage(&dto)
}

func (lm *LogMe) SetTimestampLayout(timestampLayout string) {
	lm.messageProducer.SetTimestampLayout(timestampLayout)
}

func getTimestamp() time.Time {
	return time.Now()
}

func (lm *LogMe) SetLevel(level levels.LogLevel) {
	lm.messageProducer.SetLevel(level)
}

func (lm *LogMe) AddOutput(output entity.LogOutput) {
	lm.messageProducer.AddOutput(output)
}

func GetRuntimeMetrics() entity.DebugInfo {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return entity.DebugInfo{
		NumCPU:       runtime.NumCPU(),
		CgoCalls:     int(runtime.NumCgoCall()),
		NumGoroutine: runtime.NumGoroutine(),
		Alloc:        int(m.Alloc / 1024 / 1024),
		TotalAlloc:   int(m.TotalAlloc / 1024 / 1024),
		Sys:          int(m.Sys / 1024 / 1024),
		NumGC:        int(m.NumGC),
	}
}
