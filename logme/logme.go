package logme

import (
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/dto"
	"github.com/rockkley/logme/logme/entity/levels"
	"github.com/rockkley/logme/logme/outputs"
	"time"
)

type LogMe struct {
	outputs         []outputs.LogOutput
	messageProducer *entity.MessageProducer
}

func NewLogMe() *LogMe {
	return &LogMe{
		messageProducer: entity.NewMessageProducer(),
	}
}

func (lm *LogMe) Warning(message string) {
	ts := getTimestamp()
	level := levels.Warning
	if !lm.messageProducer.Validate(message, level) {
		return
	}

	dtoMsg := dto.MessageDTO{
		Level:     level,
		Text:      message,
		Timestamp: ts,
	}
	msg := lm.messageProducer.NewMessage(&dtoMsg)
	lm.sendToOutputs(msg)
}

// Calls by level

func (lm *LogMe) sendToOutputs(message *entity.Message) {
	for _, o := range lm.outputs {
		if err := o.Write(message); err != nil { // TODO запускать в горутинах
			return // TODO не упускать ошибку, попробовать через панику
		}
	}
}

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

func (lm *LogMe) AddOutput(output outputs.LogOutput) {
	lm.outputs = append(lm.outputs, output)
}
