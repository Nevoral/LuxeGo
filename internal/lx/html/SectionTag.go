package html

import (
	"LuxeGo/internal/lx"
)

// Section -
func Section(tags ...lx.Content) *SectionTag {
	return &SectionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "section", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SectionTag struct {
	*ComponentHtmlTag
}
