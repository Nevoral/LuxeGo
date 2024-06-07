package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Section -
func Section(tags ...LuxeGo.Content) *SectionTag {
	return &SectionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "section", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SectionTag struct {
	*ComponentHtmlTag
}
