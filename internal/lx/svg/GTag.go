package svg

import (
	"LuxeGo/internal/lx"
)

// G -
func G(tags ...lx.Content) *GTag {
	return &GTag{ComponentSvgTag: &ComponentSvgTag{Name: "g", Attributes: &lx.Attributes{}, Children: &tags}}
}

type GTag struct {
	*ComponentSvgTag
}
