package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Slot -
func Slot(tags ...interface{}) *SlotTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &SlotTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "slot", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SlotTag struct {
	*ComponentHtmlTag
}

// Name -
func (s *SlotTag) Name(value string) *SlotTag {
	s.AddAttribute("name", value)
	return s
}
