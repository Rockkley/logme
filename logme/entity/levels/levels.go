package levels

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

//func (ll *LogLevel) Warning(message *entity.Message) {
//	if *ll < Warning {
//		return
//	}
//	//formattedMessage := lm.formatMessage(message)
//
//}

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
