package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Svg -
func Svg(tags ...interface{}) *SvgTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &SvgTag{ComponentSvgTag: &ComponentSvgTag{Name: "svg", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SvgTag struct {
	*ComponentSvgTag
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

// ViewBox -
func (s *SvgTag) ViewBox(value string) *SvgTag {
	s.AddAttribute("viewbox", value)
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

// Xmlns -
func (s *SvgTag) Xmlns(value string) *SvgTag {
	s.AddAttribute("xmlns", value)
	return s
}

// XmlnsXlink -
func (s *SvgTag) XmlnsXlink(value string) *SvgTag {
	s.AddAttribute("xmlns:xlink", value)
	return s
}
