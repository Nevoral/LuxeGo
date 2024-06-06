package svg

import (
	"LuxeGo/internal/lx"
)

// Script -
func Script(tags ...lx.Content) *ScriptTag {
	return &ScriptTag{ComponentSvgTag: &ComponentSvgTag{Name: "script", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ScriptTag struct {
	*ComponentSvgTag
}

// Href -
func (s *ScriptTag) Href(value string) *ScriptTag {
	s.AddAttribute("href", value)
	return s
}

// Type -
func (s *ScriptTag) Type(value string) *ScriptTag {
	s.AddAttribute("type", value)
	return s
}
