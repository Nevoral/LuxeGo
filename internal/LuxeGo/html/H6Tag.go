package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// H6 -
func H6(tags ...LuxeGo.Content) *H6Tag {
	return &H6Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h6", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H6Tag struct {
	*ComponentHtmlTag
}
