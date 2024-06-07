package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Dt -
func Dt(tags ...LuxeGo.Content) *DtTag {
	return &DtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dt", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DtTag struct {
	*ComponentHtmlTag
}
