package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Html -
func Html(tags ...LuxeGo.Content) *HtmlTag {
	return &HtmlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "html", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type HtmlTag struct {
	*ComponentHtmlTag
}

// Xmlns -
func (h *HtmlTag) Xmlns(value string) *HtmlTag {
	h.AddAttribute("xmlns", value)
	return h
}
