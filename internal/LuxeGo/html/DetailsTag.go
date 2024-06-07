package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Details -
func Details(tags ...LuxeGo.Content) *DetailsTag {
	return &DetailsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "details", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DetailsTag struct {
	*ComponentHtmlTag
}

// Open -
func (d *DetailsTag) Open() *DetailsTag {
	d.AddAttribute("open", "")
	return d
}
