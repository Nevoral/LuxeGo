package svg

import (
	"LuxeGo/internal/lx"
)

//Hatchpath - 
func Hatchpath(tags ...lx.Content) *HatchpathTag {
	return &HatchpathTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hatchpath", Attributes: &lx.Attributes{}, Children: &tags}}
}

type HatchpathTag struct {
	*ComponentHtmlTag
}

//D - 
func (h *HatchpathTag) D(value string) *HatchpathTag {
	h.AddAttribute("d", value)
	return h
}

//Offset - 
func (h *HatchpathTag) Offset(value string) *HatchpathTag {
	h.AddAttribute("offset", value)
	return h
}

//StrokeWidth - 
func (h *HatchpathTag) StrokeWidth(value string) *HatchpathTag {
	h.AddAttribute("strokewidth", value)
	return h
}

//PathLength - 
func (h *HatchpathTag) PathLength(value string) *HatchpathTag {
	h.AddAttribute("pathlength", value)
	return h
}
