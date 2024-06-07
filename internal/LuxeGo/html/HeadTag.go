package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Head -
func Head(tags ...LuxeGo.Content) *HeadTag {
	return &HeadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "head", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type HeadTag struct {
	*ComponentHtmlTag
}
