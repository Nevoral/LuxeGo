package html

import (
	"LuxeGo/internal/lx"
)

// Samp -
func Samp(tags ...lx.Content) *SampTag {
	return &SampTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "samp", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SampTag struct {
	*ComponentHtmlTag
}
