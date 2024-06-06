package svg

import (
	"LuxeGo/internal/lx"
)

// Mpath -
func Mpath(tags ...lx.Content) *MpathTag {
	return &MpathTag{ComponentSvgTag: &ComponentSvgTag{Name: "mpath", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MpathTag struct {
	*ComponentSvgTag
}

// Href -
func (m *MpathTag) Href(value string) *MpathTag {
	m.AddAttribute("href", value)
	return m
}
