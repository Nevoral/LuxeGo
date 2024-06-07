package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Code -
func Code(tags ...LuxeGo.Content) *CodeTag {
	return &CodeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "code", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type CodeTag struct {
	*ComponentHtmlTag
}
