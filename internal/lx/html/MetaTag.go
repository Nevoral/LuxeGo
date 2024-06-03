package html

import (
	"LuxeGo/internal/lx"
)

//Meta - 
func Meta() *MetaTag {
	return &MetaTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "meta", Attributes: &lx.Attributes{}, Children: nil}}}
}

type MetaTag struct {
	*ComponentTag
}

//Charset - Specifies the character encoding for the HTML document.
func (m *MetaTag) Charset(value string) *MetaTag {
	m.AddAttribute("charset", value)
	return m
}

//Content - 
func (m *MetaTag) Content(value string) *MetaTag {
	m.AddAttribute("content", value)
	return m
}

//Http-equiv - 
func (m *MetaTag) Http-equiv(value string) *MetaTag {
	m.AddAttribute("http-equiv", value)
	return m
}

//Name - 
func (m *MetaTag) Name(value string) *MetaTag {
	m.AddAttribute("name", value)
	return m
}
