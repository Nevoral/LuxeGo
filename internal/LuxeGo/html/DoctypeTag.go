package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// DOCTYPE -
func DOCTYPE() *DoctypeTag {
	return &DoctypeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!DOCTYPE", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type DoctypeTag struct {
	*ComponentHtmlTag
}

// Html -
func (d *DoctypeTag) Html() *DoctypeTag {
	d.AddAttribute("html", "")
	return d
}
