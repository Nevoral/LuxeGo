package html

import (
	"LuxeGo/internal/lx"
)

// Code -
func Code(tags ...lx.Content) *CodeTag {
	return &CodeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "code", Attributes: &lx.Attributes{}, Children: &tags}}
}

type CodeTag struct {
	*ComponentHtmlTag
}
