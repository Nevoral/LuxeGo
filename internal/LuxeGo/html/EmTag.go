package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Em -
func Em(tags ...LuxeGo.Content) *EmTag {
	return &EmTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "em", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type EmTag struct {
	*ComponentHtmlTag
}
