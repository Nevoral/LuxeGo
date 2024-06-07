package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Thead -
func Thead(tags ...LuxeGo.Content) *TheadTag {
	return &TheadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "thead", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TheadTag struct {
	*ComponentHtmlTag
}
