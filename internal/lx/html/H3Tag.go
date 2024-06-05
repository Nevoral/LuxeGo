package html

import (
	"LuxeGo/internal/lx"
)

//H3 - 
func H3(tags ...lx.Content) *H3Tag {
	return &H3Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h3", Attributes: &lx.Attributes{}, Children: &tags}}
}

type H3Tag struct {
	*ComponentHtmlTag
}
