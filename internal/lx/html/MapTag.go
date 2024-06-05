package html

import (
	"LuxeGo/internal/lx"
)

//Map - 
func Map(tags ...lx.Content) *MapTag {
	return &MapTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "map", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MapTag struct {
	*ComponentHtmlTag
}

//Name - 
func (m *MapTag) Name(value string) *MapTag {
	m.AddAttribute("name", value)
	return m
}
