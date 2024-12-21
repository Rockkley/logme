package console_output

import (
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/levels"
	"github.com/rockkley/logme/logme/outputs/console_output/visual"
	"strings"
)

const defaultFormatString = "{BackgroundColor}{Timestamp}{ColorReset}{TextColor} {TextStyle}{Level}: {Text}{ColorReset}"

var levelToString = map[levels.LogLevel]string{
	levels.Info:     "INFO",
	levels.Warning:  "WARNING",
	levels.Debug:    "DEBUG",
	levels.Critical: "CRITICAL",
}

type ConsoleOutput struct {
	FormatString string
	LevelDesigns map[levels.LogLevel]visual.MessageDesign
}

type ConsoleMessage struct {
	Timestamp string
	Level     string
	Text      string
}

func NewConsoleOutput() *ConsoleOutput {
	levelDesigns := make(map[levels.LogLevel]visual.MessageDesign)
	for level := range levelToString {
		levelDesigns[level] = getDesignForLevel(level)
	}
	consoleOutput := ConsoleOutput{
		FormatString: defaultFormatString,
		LevelDesigns: levelDesigns,
	}
	return &consoleOutput
}

func (c *ConsoleOutput) Write(message *entity.Message) (err error) {
	consoleMessage := convertToConsoleMessage(*message)
	_, err = fmt.Println(consoleMessage)
	return
}

func convertToConsoleMessage(message entity.Message) *ConsoleMessage {
	// TODO ...
	//consoleMessage = ConsoleMessage{
	//	Design:    c.LevelDesigns[message.Level],
	//	Timestamp: message.Timestamp,
	//	Level:     levelToString[message.Level],
	//	Text:      strings.ToLower(message.Text),
	//}
	//
	//formatData := map[string]string{
	//	"BackgroundColor": consoleMessage.Design.ColorPalette.BackgroundColor,
	//	"Timestamp":       message.Timestamp,
	//	"ColorReset":      visual.ColorReset,
	//	"TextColor":       consoleMessage.Design.ColorPalette.TextColor,
	//	"TextStyle":       consoleMessage.Design.TextStyle,
	//	"Level":           levelToString[message.Level],
	//	"Text":            strings.ToLower(message.Text),
	//}
	//
	//formatString := c.FormatString
	//if formatString == "" {
	//	formatString = defaultFormatString
	//}
	//
	//out := mapToFormatString(formatData, formatString)
	return &ConsoleMessage{}
}
func getDesignForLevel(level levels.LogLevel) visual.MessageDesign {
	switch level {
	case levels.Info:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorGreen, BackgroundColor: visual.BgGreen},
			TextStyle:    "",
		}
	case levels.Warning:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorYellow, BackgroundColor: visual.BgYellow},
			TextStyle:    "",
		}
	case levels.Debug:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorBlue, BackgroundColor: visual.BgBlue},
			TextStyle:    visual.ItalicText,
		}
	case levels.Critical:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorRed, BackgroundColor: visual.BgRed},
			TextStyle:    visual.BoldText,
		}
	default:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorWhite, BackgroundColor: visual.BgBlue},
			TextStyle:    "",
		}
	}
}

func mapToFormatString(data map[string]string, format string) string {
	var builder strings.Builder
	lastIndex := 0

	for lastIndex < len(format) {
		nextIndex := strings.IndexByte(format[lastIndex:], '{')
		if nextIndex == -1 {
			builder.WriteString(format[lastIndex:])
			break
		}
		nextIndex += lastIndex

		builder.WriteString(format[lastIndex:nextIndex])
		endIndex := strings.IndexByte(format[nextIndex:], '}')
		if endIndex == -1 {
			builder.WriteString(format[nextIndex:])
			break
		}
		endIndex += nextIndex

		key := format[nextIndex+1 : endIndex]
		if value, ok := data[key]; ok {
			builder.WriteString(value)
		} else {
			builder.WriteString(format[nextIndex : endIndex+1])
		}
		lastIndex = endIndex + 1
	}

	return builder.String()
}
