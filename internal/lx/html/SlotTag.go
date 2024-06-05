package html

import (
	"LuxeGo/internal/lx"
)

//Slot - 
func Slot(tags ...lx.Content) *SlotTag {
	return &SlotTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "slot", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SlotTag struct {
	*ComponentHtmlTag
}

//Name - 
func (s *SlotTag) Name(value string) *SlotTag {
	s.AddAttribute("name", value)
	return s
}
