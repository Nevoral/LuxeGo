package html

import (
	"LuxeGo/internal/lx"
)

//Source - 
func Source() *SourceTag {
	return &SourceTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "source", Attributes: &lx.Attributes{}, Children: nil}}
}

type SourceTag struct {
	*ComponentHtmlTag
}

//Type - Specifies the type of an <input> element.
func (s *SourceTag) Type(value string) *SourceTag {
	s.AddAttribute("type", value)
	return s
}

//Src - Specifies the URL of an image.
func (s *SourceTag) Src(value string) *SourceTag {
	s.AddAttribute("src", value)
	return s
}

//Srcset - 
func (s *SourceTag) Srcset(value string) *SourceTag {
	s.AddAttribute("srcset", value)
	return s
}

//Sizes - 
func (s *SourceTag) Sizes(value string) *SourceTag {
	s.AddAttribute("sizes", value)
	return s
}

//Media - 
func (s *SourceTag) Media(value string) *SourceTag {
	s.AddAttribute("media", value)
	return s
}

//Height - 
func (s *SourceTag) Height(value string) *SourceTag {
	s.AddAttribute("height", value)
	return s
}

//Width - 
func (s *SourceTag) Width(value string) *SourceTag {
	s.AddAttribute("width", value)
	return s
}
