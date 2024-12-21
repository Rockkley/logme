package dto

import (
	"github.com/rockkley/logme/logme/entity/levels"
	"time"
)

type MessageDTO struct {
	Level     levels.LogLevel
	Text      string
	Timestamp time.Time
}

//func (m *MessageDTO) Validate(){
//	if m.Text == ""  {
//		return
//	}
//}