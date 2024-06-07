package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Svg -
func Svg(tags ...LuxeGo.Content) *SvgTag {
	return &SvgTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "svg", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SvgTag struct {
	*ComponentHtmlTag
}
