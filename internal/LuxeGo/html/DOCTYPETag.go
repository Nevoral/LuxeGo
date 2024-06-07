package html

import (
	"LuxeGo/internal/LuxeGo"
)

// DOCTYPE -
func DOCTYPE() *DOCTYPETag {
	return &DOCTYPETag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!DOCTYPE", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type DOCTYPETag struct {
	*ComponentHtmlTag
}

// Html -
func (d *DOCTYPETag) Html() *DOCTYPETag {
	d.AddAttribute("html", "")
	return d
}
