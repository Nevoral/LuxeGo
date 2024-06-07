package html

import (
	"github.com/Nevoral/LuxeGo"
)

// H3 -
func H3(tags ...LuxeGo.Content) *H3Tag {
	return &H3Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h3", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H3Tag struct {
	*ComponentHtmlTag
}
