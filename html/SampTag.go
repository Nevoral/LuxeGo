package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Samp -
func Samp(tags ...LuxeGo.Content) *SampTag {
	return &SampTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "samp", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SampTag struct {
	*ComponentHtmlTag
}
