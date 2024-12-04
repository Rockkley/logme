package visual

import (
	"reflect"
)

type ColorPalette struct {
	TextColor       string
	BackgroundColor string
}

type MessageDesign struct {
	ColorPalette ColorPalette
	TextStyle    string
}

func (md *MessageDesign) GetSize() int {
	return int(reflect.TypeOf(md).Size())
}
