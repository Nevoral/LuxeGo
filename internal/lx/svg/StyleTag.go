package svg

import (
	"LuxeGo/internal/lx"
)

// Style -
func Style(tags ...lx.Content) *StyleTag {
	return &StyleTag{ComponentSvgTag: &ComponentSvgTag{Name: "style", Attributes: &lx.Attributes{}, Children: &tags}}
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
