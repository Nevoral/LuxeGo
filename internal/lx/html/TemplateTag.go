package html

import (
	"LuxeGo/internal/lx"
)

// Template -
func Template(tags ...lx.Content) *TemplateTag {
	return &TemplateTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "template", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TemplateTag struct {
	*ComponentHtmlTag
}
