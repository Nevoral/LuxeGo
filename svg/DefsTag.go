package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Defs -
func Defs(tags ...LuxeGo.Content) *DefsTag {
	return &DefsTag{ComponentSvgTag: &ComponentSvgTag{Name: "defs", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DefsTag struct {
	*ComponentSvgTag
}
