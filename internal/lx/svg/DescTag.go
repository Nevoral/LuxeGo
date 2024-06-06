package svg

import (
	"LuxeGo/internal/lx"
)

// Desc -
func Desc(tags ...lx.Content) *DescTag {
	return &DescTag{ComponentSvgTag: &ComponentSvgTag{Name: "desc", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DescTag struct {
	*ComponentSvgTag
}
