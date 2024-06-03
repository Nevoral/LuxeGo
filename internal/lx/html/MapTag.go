package html

import (
	"LuxeGo/internal/lx"
)

//Map - 
func Map(tags ...lx.Content) *MapTag {
	return &MapTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "map", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type MapTag struct {
	*ComponentTag
}

//Name - 
func (m *MapTag) Name(value string) *MapTag {
	m.AddAttribute("name", value)
	return m
}
