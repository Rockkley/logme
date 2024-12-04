package logme

import (
	"github.com/rockkley/logme/logme/entity"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Debug
	Critical
	All
)

//func (lm *LogMe) Info(str string) {
//	if lm.level < Info {
//		return
//	}
//	out := fmt.Sprintf("%s%s%s%s %s: %s%s", visual2.BgGreen, lm.getTimestamp(), visual2.ColorReset, visual2.ColorGreen, "INFO", strings.ToLower(str), visual2.ColorReset)
//}

func (lm *LogMe) Warning(str string) {
	if lm.level < Warning {
		return
	}
	message := entity.NewMessage(int(Warning), str, lm.timestampLayout)
	//formattedMessage := lm.formatMessage(message)
	for _, o := range lm.outputs {
		if err := o.Write(message); err != nil {
			return // TODO не упускать ошибку
		}
	}
}

//func (lm *LogMe) Critical(str string) {
//	if lm.level < Critical {
//		return
//	}
//	out := fmt.Sprintf("%s%s%s%s %s%s: %s%s%s", visual2.BgRed, lm.getTimestamp(), visual2.ColorReset, visual2.ColorRed, visual2.BoldText, "CRITICAL", strings.ToLower(str), visual2.ColorReset, visual2.ColorReset)
//}

//func (lm *LogMe) Debug(str string) {
//	if lm.level < Debug {
//		return
//	}
//
//	out := fmt.Sprintf("%s%s%s%s %s%s: %s%s%s", visual2.BgBlue, lm.getTimestamp(), visual2.ColorReset, visual2.ColorBlue, visual2.ItalicText, "DEBUG", strings.ToLower(str), visual2.ColorReset, visual2.ColorReset)
//}
//
//func (lm *LogMe) formatMessage(msg *entity.Message) string {
//	return fmt.Sprintf("[30%s] %d: %s", msg.TimeStamp, msg.Level, msg.Text)
//}
