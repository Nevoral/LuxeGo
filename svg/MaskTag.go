package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Mask -
func Mask(tags ...interface{}) *MaskTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &MaskTag{ComponentSvgTag: &ComponentSvgTag{Name: "mask", Attributes: &LuxeGo.Attributes{}, Children: &children}}
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
