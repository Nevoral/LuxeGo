package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Hr -
func Hr() *HrTag {
	return &HrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hr", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type HrTag struct {
	*ComponentHtmlTag
}
