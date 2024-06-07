package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// H2 -
func H2(tags ...LuxeGo.Content) *H2Tag {
	return &H2Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h2", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H2Tag struct {
	*ComponentHtmlTag
}
