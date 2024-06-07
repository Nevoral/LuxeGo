package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// FeTile -
func FeTile() *FeTileTag {
	return &FeTileTag{ComponentSvgTag: &ComponentSvgTag{Name: "fetile", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeTileTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeTileTag) In(value string) *FeTileTag {
	f.AddAttribute("in", value)
	return f
}
