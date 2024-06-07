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

// Width -
func (s *SvgTag) Width(value string) *SvgTag {
	s.AddAttribute("width", value)
	return s
}

// Height -
func (s *SvgTag) Height(value string) *SvgTag {
	s.AddAttribute("height", value)
	return s
}

// PreserveAspectRatio -
func (s *SvgTag) PreserveAspectRatio(value string) *SvgTag {
	s.AddAttribute("preserveaspectratio", value)
	return s
}

// Version -
func (s *SvgTag) Version(value string) *SvgTag {
	s.AddAttribute("version", value)
	return s
}

// BaseProfile -
func (s *SvgTag) BaseProfile(value string) *SvgTag {
	s.AddAttribute("baseprofile", value)
	return s
}

// X -
func (s *SvgTag) X(value string) *SvgTag {
	s.AddAttribute("x", value)
	return s
}

// Y -
func (s *SvgTag) Y(value string) *SvgTag {
	s.AddAttribute("y", value)
	return s
}
