package svg

import (
	"LuxeGo/internal/lx"
)

//Stop - 
func Stop() *StopTag {
	return &StopTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "stop", Attributes: &lx.Attributes{}, Children: nil}}
}

type StopTag struct {
	*ComponentHtmlTag
}

//Offset - 
func (s *StopTag) Offset(value string) *StopTag {
	s.AddAttribute("offset", value)
	return s
}

//Stop-color - 
func (s *StopTag) Stop-color(value string) *StopTag {
	s.AddAttribute("stop-color", value)
	return s
}

//Stop-opacity - 
func (s *StopTag) Stop-opacity(value string) *StopTag {
	s.AddAttribute("stop-opacity", value)
	return s
}
