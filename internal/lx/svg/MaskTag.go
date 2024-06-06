package svg

import (
	"LuxeGo/internal/lx"
)

// Mask -
func Mask(tags ...lx.Content) *MaskTag {
	return &MaskTag{ComponentSvgTag: &ComponentSvgTag{Name: "mask", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MaskTag struct {
	*ComponentSvgTag
}

// Height -
func (m *MaskTag) Height(value string) *MaskTag {
	m.AddAttribute("height", value)
	return m
}

// MaskContentUnits -
func (m *MaskTag) MaskContentUnits(value string) *MaskTag {
	m.AddAttribute("maskcontentunits", value)
	return m
}

// MaskUnits -
func (m *MaskTag) MaskUnits(value string) *MaskTag {
	m.AddAttribute("maskunits", value)
	return m
}

// Width -
func (m *MaskTag) Width(value string) *MaskTag {
	m.AddAttribute("width", value)
	return m
}

// X -
func (m *MaskTag) X(value string) *MaskTag {
	m.AddAttribute("x", value)
	return m
}

// Y -
func (m *MaskTag) Y(value string) *MaskTag {
	m.AddAttribute("y", value)
	return m
}
