package html

import (
	"LuxeGo/internal/lx"
)

// DOCTYPE -
func DOCTYPE() *DOCTYPETag {
	return &DOCTYPETag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!DOCTYPE", Attributes: &lx.Attributes{}, Children: nil}}
}

type DOCTYPETag struct {
	*ComponentHtmlTag
}

// Html -
func (d *DOCTYPETag) Html() *DOCTYPETag {
	d.AddAttribute("html", "")
	return d
}
