package html

import (
	"LuxeGo/internal/lx"
)

// Head -
func Head(tags ...lx.Content) *HeadTag {
	return &HeadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "head", Attributes: &lx.Attributes{}, Children: &tags}}
}

type HeadTag struct {
	*ComponentHtmlTag
}
