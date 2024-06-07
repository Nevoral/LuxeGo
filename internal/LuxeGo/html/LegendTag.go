package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Legend -
func Legend(tags ...LuxeGo.Content) *LegendTag {
	return &LegendTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "legend", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type LegendTag struct {
	*ComponentHtmlTag
}