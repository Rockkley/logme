package logme

import (
	outputs "github.com/rockkley/logme/logme/outputs"
	"time"
)

const (
	defaultTimestampLayout = time.DateTime
)

type LogMe struct {
	timestampLayout string
	level           LogLevel
	outputs         []outputs.LogOutput
}

func NewLogMe() *LogMe {
	return &LogMe{
		timestampLayout: defaultTimestampLayout,
		level:           All,
	}
}

// Timestamp

func (lm *LogMe) SetTimestampLayout(timestampLayout string) {
	lm.timestampLayout = timestampLayout
}

func (lm *LogMe) getTimestamp() string {
	return time.Now().Format(lm.timestampLayout)
}

func (lm *LogMe) SetLevel(level LogLevel) {
	lm.level = level
}

func (lm *LogMe) AddOutput(output outputs.LogOutput) {
	lm.outputs = append(lm.outputs, output)
}
