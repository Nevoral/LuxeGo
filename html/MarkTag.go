package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Mark -
func Mark(tags ...LuxeGo.Content) *MarkTag {
	return &MarkTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "mark", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MarkTag struct {
	*ComponentHtmlTag
}
