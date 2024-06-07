package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Hatch -
func Hatch(tags ...interface{}) *HatchTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &HatchTag{ComponentSvgTag: &ComponentSvgTag{Name: "hatch", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HatchTag struct {
	*ComponentSvgTag
}

// X -
func (h *HatchTag) X(value string) *HatchTag {
	h.AddAttribute("x", value)
	return h
}

// Y -
func (h *HatchTag) Y(value string) *HatchTag {
	h.AddAttribute("y", value)
	return h
}

// Pitch -
func (h *HatchTag) Pitch(value string) *HatchTag {
	h.AddAttribute("pitch", value)
	return h
}

// Rotate -
func (h *HatchTag) Rotate(value string) *HatchTag {
	h.AddAttribute("rotate", value)
	return h
}

// HatchUnits -
func (h *HatchTag) HatchUnits(value string) *HatchTag {
	h.AddAttribute("hatchunits", value)
	return h
}

// HatchContentUnits -
func (h *HatchTag) HatchContentUnits(value string) *HatchTag {
	h.AddAttribute("hatchcontentunits", value)
	return h
}

// Transform -
func (h *HatchTag) Transform(value string) *HatchTag {
	h.AddAttribute("transform", value)
	return h
}
