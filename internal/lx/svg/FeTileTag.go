package svg

import (
	"LuxeGo/internal/lx"
)

//FeTile - 
func FeTile() *FeTileTag {
	return &FeTileTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fetile", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeTileTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeTileTag) In(value string) *FeTileTag {
	f.AddAttribute("in", value)
	return f
}
