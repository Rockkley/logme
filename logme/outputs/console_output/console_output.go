package console_output

import (
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/levels"
	"github.com/rockkley/logme/logme/outputs/console_output/visual"
	"strings"
	"time"
)

type levelDesignsCache = map[levels.LogLevel]visual.MessageDesign

const defaultFormatString = "{BackgroundColor}{Timestamp}{ColorReset}{TextColor} {TextStyle}{Level}: {Text}{ColorReset}"

var levelToString = map[levels.LogLevel]string{
	levels.Info:     "INFO",
	levels.Warning:  "WARNING",
	levels.Debug:    "DEBUG",
	levels.Critical: "CRITICAL",
}

type ConsoleOutput struct {
	FormatString string
	LevelDesigns levelDesignsCache
}

func NewConsoleOutput() *ConsoleOutput {
	levelDesigns := createLevelDesignsCache()

	consoleOutput := ConsoleOutput{
		FormatString: defaultFormatString,
		LevelDesigns: levelDesigns,
	}
	return &consoleOutput
}

func (c *ConsoleOutput) Write(message *entity.Message) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic in console output Write")
		}
	}()
	consoleMessage := c.convertToConsoleMessage(*message)
	if _, err := fmt.Println(consoleMessage); err != nil {
		panic(err)
	}
	return
}

func createLevelDesignsCache() levelDesignsCache {
	cache := make(levelDesignsCache)
	// INFO
	cache[levels.Info] = visual.MessageDesign{
		ColorPalette: visual.ColorPalette{TextColor: visual.ColorGreen, BackgroundColor: visual.BgGreen},
		TextStyle:    "",
	}
	// WARNING
	cache[levels.Warning] = visual.MessageDesign{
		ColorPalette: visual.ColorPalette{TextColor: visual.ColorYellow, BackgroundColor: visual.BgYellow},
		TextStyle:    "",
	}
	// DEBUG
	cache[levels.Debug] = visual.MessageDesign{
		ColorPalette: visual.ColorPalette{TextColor: visual.ColorBlue, BackgroundColor: visual.BgBlue},
		TextStyle:    visual.ItalicText,
	}
	// CRITICAL
	cache[levels.Critical] = visual.MessageDesign{
		ColorPalette: visual.ColorPalette{TextColor: visual.ColorRed, BackgroundColor: visual.BgRed},
		TextStyle:    visual.BoldText,
	}
	return cache
}

func (c *ConsoleOutput) convertToConsoleMessage(message entity.Message) string {
	design := c.LevelDesigns[message.Level]

	formatData := map[string]string{
		"BackgroundColor": design.ColorPalette.BackgroundColor,
		"Timestamp":       message.Timestamp.Format(time.DateTime),
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

	out := mapToFormatString(formatData, formatString)
	return out
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
