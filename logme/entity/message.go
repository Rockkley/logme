package entity

import (
	"bytes"
	"github.com/rockkley/logme/logme/entity/levels"
)

type Message struct {
	Level     levels.LogLevel
	Text      bytes.Buffer
	Timestamp string
}
