package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Pre -
func Pre(tags ...LuxeGo.Content) *PreTag {
	return &PreTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "pre", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type PreTag struct {
	*ComponentHtmlTag
}
