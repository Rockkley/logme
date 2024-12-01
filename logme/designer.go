package logme

import (
	"github.com/rockkley/logme/logme/visual"
	"strings"
)

func DesignMessage(prop *visual.MessageDesign, message string) {
	if message == "" || prop == nil {

		return
	}

	messageSize := prop.GetSize()

	var output strings.Builder

	output.Grow(messageSize + len(message))

	if prop.Color != "" {
		output.WriteString(prop.Color)
	}
	output.WriteString("d")

}
