package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Sub -
func Sub(tags ...LuxeGo.Content) *SubTag {
	return &SubTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sub", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SubTag struct {
	*ComponentHtmlTag
}
