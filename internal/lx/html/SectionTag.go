package html

import (
	"LuxeGo/internal/lx"
)

//Section - 
func Section(tags ...lx.Content) *SectionTag {
	return &SectionTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "section", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SectionTag struct {
	*ComponentTag
}
