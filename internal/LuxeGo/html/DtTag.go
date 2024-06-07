package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Dt -
func Dt(tags ...LuxeGo.Content) *DtTag {
	return &DtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dt", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DtTag struct {
	*ComponentHtmlTag
}
