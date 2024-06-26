package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Stop -
func Stop() *StopTag {
	return &StopTag{ComponentSvgTag: &ComponentSvgTag{Name: "stop", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type StopTag struct {
	*ComponentSvgTag
}

// Offset -
func (s *StopTag) Offset(value string) *StopTag {
	s.AddAttribute("offset", value)
	return s
}

// StopColor -
func (s *StopTag) StopColor(value string) *StopTag {
	s.AddAttribute("stop-color", value)
	return s
}

// StopOpacity -
func (s *StopTag) StopOpacity(value string) *StopTag {
	s.AddAttribute("stop-opacity", value)
	return s
}
