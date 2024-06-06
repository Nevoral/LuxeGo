package html

import (
	"LuxeGo/internal/lx"
)

// Kbd -
func Kbd(tags ...lx.Content) *KbdTag {
	return &KbdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "kbd", Attributes: &lx.Attributes{}, Children: &tags}}
}

type KbdTag struct {
	*ComponentHtmlTag
}
