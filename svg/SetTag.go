package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Set -
func Set(tags ...LuxeGo.Content) *SetTag {
	return &SetTag{ComponentSvgTag: &ComponentSvgTag{Name: "set", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SetTag struct {
	*ComponentSvgTag
}

// To -
func (s *SetTag) To(value string) *SetTag {
	s.AddAttribute("to", value)
	return s
}

// AttributeName -
func (s *SetTag) AttributeName(value string) *SetTag {
	s.AddAttribute("attributename", value)
	return s
}

// AttributeType -
func (s *SetTag) AttributeType(value string) *SetTag {
	s.AddAttribute("attributetype", value)
	return s
}

// Begin -
func (s *SetTag) Begin(value string) *SetTag {
	s.AddAttribute("begin", value)
	return s
}

// Dur -
func (s *SetTag) Dur(value string) *SetTag {
	s.AddAttribute("dur", value)
	return s
}

// End -
func (s *SetTag) End(value string) *SetTag {
	s.AddAttribute("end", value)
	return s
}

// Fill -
func (s *SetTag) Fill(value string) *SetTag {
	s.AddAttribute("fill", value)
	return s
}

// RepeatCount -
func (s *SetTag) RepeatCount(value string) *SetTag {
	s.AddAttribute("repeatcount", value)
	return s
}

// RepeatDur -
func (s *SetTag) RepeatDur(value string) *SetTag {
	s.AddAttribute("repeatdur", value)
	return s
}

// Restart -
func (s *SetTag) Restart(value string) *SetTag {
	s.AddAttribute("restart", value)
	return s
}
