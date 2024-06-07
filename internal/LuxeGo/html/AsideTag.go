package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Aside -
func Aside(tags ...LuxeGo.Content) *AsideTag {
	return &AsideTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "aside", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type AsideTag struct {
	*ComponentHtmlTag
}
