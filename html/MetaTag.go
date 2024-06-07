package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Meta -
func Meta() *MetaTag {
	return &MetaTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "meta", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type MetaTag struct {
	*ComponentHtmlTag
}

// Charset - Specifies the character encoding for the HTML document.
func (m *MetaTag) Charset(value string) *MetaTag {
	m.AddAttribute("charset", value)
	return m
}

// Content -
func (m *MetaTag) Content(value string) *MetaTag {
	m.AddAttribute("content", value)
	return m
}

// HttpEquiv -
func (m *MetaTag) HttpEquiv(value string) *MetaTag {
	m.AddAttribute("http-equiv", value)
	return m
}

// Name -
func (m *MetaTag) Name(value string) *MetaTag {
	m.AddAttribute("name", value)
	return m
}
