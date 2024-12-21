package visual

import (
	"reflect"
)

type MessageDesign struct {
	ColorPalette ColorPalette
	TextStyle    string
}

type ColorPalette struct {
	TextColor       string
	BackgroundColor string
}

func (md *MessageDesign) GetSize() int {
	return int(reflect.TypeOf(md).Size())
}
