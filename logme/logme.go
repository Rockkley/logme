package logme

import (
	"github.com/rockkley/logme/logme/entity"
	dto2 "github.com/rockkley/logme/logme/entity/dto"
	"github.com/rockkley/logme/logme/entity/levels"
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
	dto := dto2.MessageDTO{
		Level:     levels.Debug,
		Text:      message,
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
