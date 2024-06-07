package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Mpath -
func Mpath(tags ...LuxeGo.Content) *MpathTag {
	return &MpathTag{ComponentSvgTag: &ComponentSvgTag{Name: "mpath", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MpathTag struct {
	*ComponentSvgTag
}

// Href -
func (m *MpathTag) Href(value string) *MpathTag {
	m.AddAttribute("href", value)
	return m
}
