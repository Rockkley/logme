package entity

import (
	"github.com/rockkley/logme/logme/entity/dto"
	"github.com/rockkley/logme/logme/entity/levels"
	"strings"
	"time"
)

const (
	defaultTimestampLayout = time.DateTime
)

type MessageProducer struct {
	level           levels.LogLevel
	timestampLayout string
	outputs         []LogOutput
}

func NewMessageProducer() *MessageProducer {
	return &MessageProducer{timestampLayout: defaultTimestampLayout, level: levels.All}
}

func (mp *MessageProducer) Info(message string, ts time.Time) {
	level := levels.Info
	if !mp.Validate(message, level) {
		return
	}
	dtoMsg := dto.MessageDTO{
		Level:     level,
		Text:      message,
		Timestamp: ts,
	}
	msg := mp.newMessage(&dtoMsg)
	mp.sendToOutputs(msg)
}

func (mp *MessageProducer) Debug(message string, ts time.Time) {
	level := levels.Debug
	if !mp.Validate(message, level) {
		return
	}
	dtoMsg := dto.MessageDTO{
		Level:     level,
		Text:      message,
		Timestamp: ts,
	}
	msg := mp.newMessage(&dtoMsg)
	mp.sendToOutputs(msg)
}

func (mp *MessageProducer) Warning(message string, ts time.Time) {
	level := levels.Warning
	if !mp.Validate(message, level) {
		return
	}
	dtoMsg := dto.MessageDTO{
		Level:     level,
		Text:      message,
		Timestamp: ts,
	}
	msg := mp.newMessage(&dtoMsg)
	mp.sendToOutputs(msg)
}

func (mp *MessageProducer) Critical(message string, ts time.Time) {
	level := levels.Critical
	if !mp.Validate(message, level) {
		return
	}
	dtoMsg := dto.MessageDTO{
		Level:     level,
		Text:      message,
		Timestamp: ts,
	}
	msg := mp.newMessage(&dtoMsg)
	mp.sendToOutputs(msg)
}

func (mp *MessageProducer) newMessage(dto *dto.MessageDTO) (msg *Message) {
	msg = &Message{
		Level:     dto.Level,
		Text:      dto.Text,
		Timestamp: dto.Timestamp,
	}
	return
}

func (mp *MessageProducer) Validate(message string, level levels.LogLevel) bool {
	if strings.TrimSpace(message) == "" || mp.level < level {
		return false
	}
	return true
}

func (mp *MessageProducer) SetLevel(level levels.LogLevel) {
	mp.level = level
}

func (mp *MessageProducer) SetTimestampLayout(timestampLayout string) {
	mp.timestampLayout = timestampLayout
}

func (mp *MessageProducer) sendToOutputs(message *Message) {
	for _, o := range mp.outputs {
		if err := o.Write(message); err != nil { // TODO запускать в горутинах
			return // TODO не упускать ошибку, попробовать через панику
		}
	}
}

func (mp *MessageProducer) AddOutput(output LogOutput) {
	mp.outputs = append(mp.outputs, output)
}
