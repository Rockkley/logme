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
}

func NewMessageProducer() *MessageProducer {
	return &MessageProducer{timestampLayout: defaultTimestampLayout, level: levels.All}
}

func (mp *MessageProducer) NewMessage(dto *dto.MessageDTO) (msg *Message) {
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
