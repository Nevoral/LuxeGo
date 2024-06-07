package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Svg -
func Svg(tags ...LuxeGo.Content) *SvgTag {
	return &SvgTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "svg", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SvgTag struct {
	*ComponentHtmlTag
}

// Xmlns -
func (s *SvgTag) Xmlns(value string) *SvgTag {
	s.AddAttribute("xmlns", value)
	return s
}

// ViewBox -
func (s *SvgTag) ViewBox(value string) *SvgTag {
	s.AddAttribute("viewbox", value)
	return s
}
