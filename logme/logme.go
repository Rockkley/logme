package logme

import (
	"fmt"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type LogMe struct {
	timestampLayout string
}

func (lm *LogMe) getTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (lm *LogMe) INFO(str string) {
	out := fmt.Sprintf("\033[42m%s\033[0m%s INFO: %s%s", lm.getTimestamp(), colorGreen, str, colorReset)
	fmt.Println(out)
}

func (lm *LogMe) WARNING(str string) {
	out := fmt.Sprintf("\033[43m%s\033[0m%s WARNING: %s%s", lm.getTimestamp(), colorYellow, str, colorReset)
	fmt.Println(out)
}

func (lm *LogMe) CRITICAL(str string) {
	out := fmt.Sprintf("\033[41m%s\033[0m%s CRITICAL: %s%s", lm.getTimestamp(), colorRed, str, colorReset)
	fmt.Println(out)
}
