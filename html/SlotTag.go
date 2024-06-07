package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Slot -
func Slot(tags ...LuxeGo.Content) *SlotTag {
	return &SlotTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "slot", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SlotTag struct {
	*ComponentHtmlTag
}

// Name -
func (s *SlotTag) Name(value string) *SlotTag {
	s.AddAttribute("name", value)
	return s
}
