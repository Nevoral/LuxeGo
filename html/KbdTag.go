package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Kbd -
func Kbd(tags ...LuxeGo.Content) *KbdTag {
	return &KbdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "kbd", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type KbdTag struct {
	*ComponentHtmlTag
}
