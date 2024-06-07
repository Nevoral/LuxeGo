package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// H1 -
func H1(tags ...LuxeGo.Content) *H1Tag {
	return &H1Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h1", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H1Tag struct {
	*ComponentHtmlTag
}
