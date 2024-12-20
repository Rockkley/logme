package entity

import (
	"github.com/rockkley/logme/logme/entity/levels"
	"time"
)

type Message struct {
	Level     levels.LogLevel
	Text      string
	Timestamp time.Time
}
