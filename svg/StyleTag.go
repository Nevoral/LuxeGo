package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Style -
func Style(tags ...LuxeGo.Content) *StyleTag {
	return &StyleTag{ComponentSvgTag: &ComponentSvgTag{Name: "style", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type StyleTag struct {
	*ComponentSvgTag
}

// Type -
func (s *StyleTag) Type(value string) *StyleTag {
	s.AddAttribute("type", value)
	return s
}

// Media -
func (s *StyleTag) Media(value string) *StyleTag {
	s.AddAttribute("media", value)
	return s
}

// Title -
func (s *StyleTag) Title(value string) *StyleTag {
	s.AddAttribute("title", value)
	return s
}
