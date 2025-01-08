package console_output

import (
	"fmt"
	"github.com/rockkley/logme/logme/entity"
	"github.com/rockkley/logme/logme/entity/levels"
	"github.com/rockkley/logme/logme/outputs/console_output/visual"
	"strings"
)

type levelDesignsCacheMap = map[levels.LogLevel]visual.MessageDesign

const defaultFormatString = "{BackgroundColor}{Timestamp}{ColorReset}{TextColor} {TextStyle}{Level}: {Text}{ColorReset}"

type ConsoleOutput struct {
	FormatString string
	LevelDesigns levelDesignsCacheMap
}

func NewConsoleOutput() *ConsoleOutput {
	levelDesignsCache := cacheLevelDesigns()

	consoleOutput := ConsoleOutput{
		FormatString: defaultFormatString,
		LevelDesigns: levelDesignsCache,
	}

	return &consoleOutput
}

func (c *ConsoleOutput) Write(message *entity.Message) error {
	consoleMessage := c.convertToConsoleMessage(message)

	if _, err := fmt.Println(consoleMessage); err != nil {
		return err
	}

	return nil
}

func cacheLevelDesigns() levelDesignsCacheMap {
	cache := make(levelDesignsCacheMap)
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

func (c *ConsoleOutput) convertToConsoleMessage(message *entity.Message) string {
	design := c.LevelDesigns[message.Level]
	formatData := map[string]string{
		"BackgroundColor": design.ColorPalette.BackgroundColor,
		"Timestamp":       message.Timestamp,
		"ColorReset":      visual.ColorReset,
		"TextColor":       design.ColorPalette.TextColor,
		"TextStyle":       design.TextStyle,
		"Level":           message.Level.String(),
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
