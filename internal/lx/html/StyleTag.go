package html

import (
	"LuxeGo/internal/lx"
)

//Style - 
func Style(tags ...lx.Content) *StyleTag {
	return &StyleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "style", Attributes: &lx.Attributes{}, Children: &tags}}
}

type StyleTag struct {
	*ComponentHtmlTag
}

//Blocking - 
func (s *StyleTag) Blocking(value string) *StyleTag {
	s.AddAttribute("blocking", value)
	return s
}

//Media - 
func (s *StyleTag) Media(value string) *StyleTag {
	s.AddAttribute("media", value)
	return s
}

//Nonce - 
func (s *StyleTag) Nonce(value string) *StyleTag {
	s.AddAttribute("nonce", value)
	return s
}
