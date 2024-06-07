package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Template -
func Template(tags ...LuxeGo.Content) *TemplateTag {
	return &TemplateTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "template", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TemplateTag struct {
	*ComponentHtmlTag
}
