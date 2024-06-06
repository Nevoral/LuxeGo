package html

import (
	"LuxeGo/internal/lx"
)

// H5 -
func H5(tags ...lx.Content) *H5Tag {
	return &H5Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h5", Attributes: &lx.Attributes{}, Children: &tags}}
}

type H5Tag struct {
	*ComponentHtmlTag
}
