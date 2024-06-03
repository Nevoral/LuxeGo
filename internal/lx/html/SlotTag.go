package html

import (
	"LuxeGo/internal/lx"
)

//Slot - 
func Slot(tags ...lx.Content) *SlotTag {
	return &SlotTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "slot", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SlotTag struct {
	*ComponentTag
}

//Name - 
func (s *SlotTag) Name(value string) *SlotTag {
	s.AddAttribute("name", value)
	return s
}
