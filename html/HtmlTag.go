package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Html -
func Html(tags ...interface{}) *HtmlTag {
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
	return &HtmlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "html", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HtmlTag struct {
	*ComponentHtmlTag
}

// Xmlns -
func (h *HtmlTag) Xmlns(value string) *HtmlTag {
	h.AddAttribute("xmlns", value)
	return h
}
