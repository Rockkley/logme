package entity

type LogOutput interface {
	Write(message *Message) error
}
