package html

import (
	"LuxeGo/internal/lx"
)

//Html - 
func Html(tags ...lx.Content) *HtmlTag {
	return &HtmlTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "html", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type HtmlTag struct {
	*ComponentTag
}

//Xmlns - 
func (h *HtmlTag) Xmlns(value string) *HtmlTag {
	h.AddAttribute("xmlns", value)
	return h
}
