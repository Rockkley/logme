package entity

import (
	"time"
)

type Message struct {
	Level           int
	Text            string
	timestampLayout string
	Timestamp       string
}

func NewMessage(level int, text string, timestampLayout string) (msg *Message) {
	msg = &Message{
		Level:           level,
		Text:            text,
		timestampLayout: timestampLayout,
	}
	msg.Timestamp = msg.getTimestamp()
	return
}

func (m *Message) getTimestamp() string {
	return time.Now().Format(m.timestampLayout)
}
