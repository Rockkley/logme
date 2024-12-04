package console_output

import (
	"fmt"
	"github.com/rockkley/logme/logme"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/outputs/console_output/visual"
	"strings"
)

type ConsoleOutput struct {
	FormatString string
}

const defaultFormatString = "{BackgroundColor}{Timestamp}{ColorReset}{TextColor} {TextStyle}{Level}: {Text}{ColorReset}"

var levelToString = map[int]string{
	int(logme.Warning): "WARNING",
	// TODO Add other levels
}

func (c *ConsoleOutput) Write(message *entity.Message) (err error) {
	design := getDesignForLevel(message.Level)

	formatData := map[string]string{
		"BackgroundColor": design.ColorPalette.BackgroundColor,
		"Timestamp":       message.Timestamp,
		"ColorReset":      visual.ColorReset,
		"TextColor":       design.ColorPalette.TextColor,
		"TextStyle":       design.TextStyle,
		"Level":           levelToString[message.Level],
		"Text":            strings.ToLower(message.Text),
	}

	formatString := c.FormatString
	if formatString == "" {
		formatString = defaultFormatString
	}

	out := autoFormat(formatData, formatString)

	_, err = fmt.Println(out)
	return
}

func getDesignForLevel(level int) visual.MessageDesign {
	switch level {
	case int(logme.Warning):
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorYellow, BackgroundColor: visual.BgYellow},
			TextStyle:    "",
		}
	default:
		return visual.MessageDesign{
			ColorPalette: visual.ColorPalette{TextColor: visual.ColorWhite, BackgroundColor: visual.BgBlue},
			TextStyle:    "",
		}
	}
}

func autoFormat(data map[string]string, format string) string {
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
