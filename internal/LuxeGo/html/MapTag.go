package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Map -
func Map(tags ...LuxeGo.Content) *MapTag {
	return &MapTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "map", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MapTag struct {
	*ComponentHtmlTag
}

// Name -
func (m *MapTag) Name(value string) *MapTag {
	m.AddAttribute("name", value)
	return m
}
