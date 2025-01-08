package entity

import "github.com/rockkley/logme/logme/entity/levels"

type Message struct {
	Level     levels.LogLevel
	Text      string
	Timestamp string
}
