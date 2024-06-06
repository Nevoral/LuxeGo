package svg

import (
	"LuxeGo/internal/lx"
)

// Title -
func Title(tags ...lx.Content) *TitleTag {
	return &TitleTag{ComponentSvgTag: &ComponentSvgTag{Name: "title", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TitleTag struct {
	*ComponentSvgTag
}
