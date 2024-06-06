package html

import (
	"LuxeGo/internal/lx"
)

// Legend -
func Legend(tags ...lx.Content) *LegendTag {
	return &LegendTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "legend", Attributes: &lx.Attributes{}, Children: &tags}}
}

type LegendTag struct {
	*ComponentHtmlTag
}
