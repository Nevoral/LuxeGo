package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Dd -
func Dd(tags ...LuxeGo.Content) *DdTag {
	return &DdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dd", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DdTag struct {
	*ComponentHtmlTag
}
