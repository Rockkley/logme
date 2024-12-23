package logme

import (
	"github.com/rockkley/logme/logme/entity"
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

//func (lm *LogMe) Info(message string) {
//	ts := getTimestamp()
//	level := levels.Info
//	if !lm.messageProducer.Validate(message, level) {
//		return
//	}
//
//	dtoMsg := dto.MessageDTO{
//		Level:     level,
//		Text:      message,
//		Timestamp: ts,
//	}
//	msg := lm.messageProducer.NewMessage(&dtoMsg)
//	lm.sendToOutputs(msg)
//}

func (lm *LogMe) Info(message string) {
	ts := getTimestamp()
	lm.messageProducer.Info(message, ts)
}

func (lm *LogMe) Warning(message string) {
	ts := getTimestamp()
	lm.messageProducer.Warning(message, ts)
}

func (lm *LogMe) Debug(message string) {
	ts := getTimestamp()
	lm.messageProducer.Debug(message, ts)
}

func (lm *LogMe) Critical(message string) {
	ts := getTimestamp()
	lm.messageProducer.Critical(message, ts)
}

// Calls by level

func (lm *LogMe) SetTimestampLayout(timestampLayout string) {
	lm.messageProducer.SetTimestampLayout(timestampLayout)
}

// Timestamp

func getTimestamp() time.Time {
	return time.Now()
}

func (lm *LogMe) SetLevel(level levels.LogLevel) {
	lm.messageProducer.SetLevel(level)
}

func (lm *LogMe) AddOutput(output entity.LogOutput) {
	lm.messageProducer.AddOutput(output)
}
