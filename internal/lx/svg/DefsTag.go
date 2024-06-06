package svg

import (
	"LuxeGo/internal/lx"
)

// Defs -
func Defs(tags ...lx.Content) *DefsTag {
	return &DefsTag{ComponentSvgTag: &ComponentSvgTag{Name: "defs", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DefsTag struct {
	*ComponentSvgTag
}
