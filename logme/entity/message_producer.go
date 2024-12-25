package entity

import (
	"github.com/rockkley/logme/logme/entity/dto"
	"github.com/rockkley/logme/logme/entity/levels"
	"strings"
	"sync"
	"time"
)

const (
	defaultTimestampLayout = time.DateTime
)

type MessageProducer struct {
	level           levels.LogLevel
	timestampLayout string
	outputs         []LogOutput
	outChan         chan *Message
}

func NewMessageProducer() *MessageProducer {
	return &MessageProducer{timestampLayout: defaultTimestampLayout, level: levels.All, outChan: make(chan *Message)}
}

func (mp *MessageProducer) NewMessage(dto *dto.MessageDTO) {
	if !mp.Validate(dto.Text, dto.Level) {
		return
	}
	msg := &Message{
		Level:     dto.Level,
		Text:      dto.Text,
		Timestamp: dto.Timestamp,
	}
	mp.sendToOutputs(msg)
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
	var wg sync.WaitGroup
	wg.Add(len(mp.outputs))
	for _, o := range mp.outputs {
		go func() {
			defer wg.Done()
			if err := o.Write(message); err != nil { // TODO запускать в горутинах
				return // TODO не упускать ошибку, попробовать через панику
			}
		}()
	}
	wg.Wait()
}

func (mp *MessageProducer) AddOutput(output LogOutput) {
	mp.outputs = append(mp.outputs, output)
}

//func (mp *MessageProducer) prepareMessage(dto *dto.MessageDTO) *Message {
//	if !mp.Validate(dto.Text, dto.Level) {
//		return nil
//	}
//	msg := mp.newMessage(dto)
//	return msg
//}
