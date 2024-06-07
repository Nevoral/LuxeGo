package html

import (
	"LuxeGo/internal/LuxeGo"
)

// P -
func P(tags ...LuxeGo.Content) *PTag {
	return &PTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "p", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type PTag struct {
	*ComponentHtmlTag
}
