package svg

import (
	"LuxeGo/internal/lx"
)

// FeTile -
func FeTile() *FeTileTag {
	return &FeTileTag{ComponentSvgTag: &ComponentSvgTag{Name: "fetile", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeTileTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeTileTag) In(value string) *FeTileTag {
	f.AddAttribute("in", value)
	return f
}
