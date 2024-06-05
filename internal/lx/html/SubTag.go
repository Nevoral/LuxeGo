package html

import (
	"LuxeGo/internal/lx"
)

//Sub - 
func Sub(tags ...lx.Content) *SubTag {
	return &SubTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sub", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SubTag struct {
	*ComponentHtmlTag
}
