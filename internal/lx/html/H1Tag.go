package html

import (
	"LuxeGo/internal/lx"
)

// H1 -
func H1(tags ...lx.Content) *H1Tag {
	return &H1Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h1", Attributes: &lx.Attributes{}, Children: &tags}}
}

type H1Tag struct {
	*ComponentHtmlTag
}
