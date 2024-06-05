package svg

import (
	"LuxeGo/internal/lx"
)

//Desc - 
func Desc(tags ...lx.Content) *DescTag {
	return &DescTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "desc", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DescTag struct {
	*ComponentHtmlTag
}
