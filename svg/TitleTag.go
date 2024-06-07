package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Title -
func Title(tags ...LuxeGo.Content) *TitleTag {
	return &TitleTag{ComponentSvgTag: &ComponentSvgTag{Name: "title", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TitleTag struct {
	*ComponentSvgTag
}
