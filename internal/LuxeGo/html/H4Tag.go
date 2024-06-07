package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// H4 -
func H4(tags ...LuxeGo.Content) *H4Tag {
	return &H4Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h4", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H4Tag struct {
	*ComponentHtmlTag
}
