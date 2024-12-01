package logme

import (
	"fmt"
	"github.com/rockkley/logme/logme/visual"
	"strings"
)

type LogLevel int

const (
	All LogLevel = iota
	Info
	Warning
	Debug
	Critical
)

func (lm *LogMe) Info(str string) {
	if lm.level != Info && lm.level != All {
		return
	}
	out := fmt.Sprintf("%s%s%s%s %s: %s%s", visual.BgGreen, lm.getTimestamp(), visual.ColorReset, visual.ColorGreen, Info, strings.ToLower(str), visual.ColorReset)
	fmt.Println(out)
}

func (lm *LogMe) Warning(str string) {
	if lm.level <= Warning {
		formattedMessage := lm.formatMessage(Warning, str)
		for _, o := range lm.outputs {
			err := o.Write(formattedMessage)
			if err != nil {
				return
			}
		}
	}
	out := fmt.Sprintf("%s%s%s%s %s: %s%s", visual.BgYellow, lm.getTimestamp(), visual.ColorReset, visual.ColorYellow, Warning, strings.ToLower(str), visual.ColorReset)
	fmt.Println(out)
}

func (lm *LogMe) Critical(str string) {
	if lm.level != Critical && lm.level != All {
		return
	}
	out := fmt.Sprintf("%s%s%s%s %s%s: %s%s%s", visual.BgRed, lm.getTimestamp(), visual.ColorReset, visual.ColorRed, visual.BoldText, Critical, strings.ToLower(str), visual.ColorReset, visual.ColorReset)
	fmt.Println(out)
}

func (lm *LogMe) Debug(str string) {
	if lm.level != Debug && lm.level != All {
		return
	}
	out := fmt.Sprintf("%s%s%s%s %s%s: %s%s%s", visual.BgBlue, lm.getTimestamp(), visual.ColorReset, visual.ColorBlue, visual.ItalicText, Debug, strings.ToLower(str), visual.ColorReset, visual.ColorReset)
	fmt.Println(out)
}

func (lm *LogMe) formatMessage(level LogLevel, message string) string {
	timestamp := lm.getTimestamp()
	return fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
}
