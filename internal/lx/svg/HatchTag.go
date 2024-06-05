package svg

import (
	"LuxeGo/internal/lx"
)

//Hatch - 
func Hatch(tags ...lx.Content) *HatchTag {
	return &HatchTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hatch", Attributes: &lx.Attributes{}, Children: &tags}}
}

type HatchTag struct {
	*ComponentHtmlTag
}

//X - 
func (h *HatchTag) X(value string) *HatchTag {
	h.AddAttribute("x", value)
	return h
}

//Y - 
func (h *HatchTag) Y(value string) *HatchTag {
	h.AddAttribute("y", value)
	return h
}

//Pitch - 
func (h *HatchTag) Pitch(value string) *HatchTag {
	h.AddAttribute("pitch", value)
	return h
}

//Rotate - 
func (h *HatchTag) Rotate(value string) *HatchTag {
	h.AddAttribute("rotate", value)
	return h
}

//HatchUnits - 
func (h *HatchTag) HatchUnits(value string) *HatchTag {
	h.AddAttribute("hatchunits", value)
	return h
}

//HatchContentUnits - 
func (h *HatchTag) HatchContentUnits(value string) *HatchTag {
	h.AddAttribute("hatchcontentunits", value)
	return h
}

//Transform - 
func (h *HatchTag) Transform(value string) *HatchTag {
	h.AddAttribute("transform", value)
	return h
}
