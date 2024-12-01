package visual

import (
	"reflect"
)

type MessageDesign struct {
	Color        string
	IsBold       bool
	IsItalic     bool
	IsUnderlined bool
}

func (md *MessageDesign) GetSize() int {
	return int(reflect.TypeOf(md).Size())
}
