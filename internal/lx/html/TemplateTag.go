package html

import (
	"LuxeGo/internal/lx"
)

//Template - 
func Template(tags ...lx.Content) *TemplateTag {
	return &TemplateTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "template", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TemplateTag struct {
	*ComponentTag
}
