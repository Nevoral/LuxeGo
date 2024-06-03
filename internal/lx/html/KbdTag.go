package html

import (
	"LuxeGo/internal/lx"
)

//Kbd - 
func Kbd(tags ...lx.Content) *KbdTag {
	return &KbdTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "kbd", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type KbdTag struct {
	*ComponentTag
}
