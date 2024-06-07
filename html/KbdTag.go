package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Kbd -
func Kbd(tags ...interface{}) *KbdTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &KbdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "kbd", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type KbdTag struct {
	*ComponentHtmlTag
}
