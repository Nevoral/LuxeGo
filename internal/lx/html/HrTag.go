package html

import (
	"LuxeGo/internal/lx"
)

//Hr - 
func Hr() *HrTag {
	return &HrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hr", Attributes: &lx.Attributes{}, Children: nil}}
}

type HrTag struct {
	*ComponentHtmlTag
}
