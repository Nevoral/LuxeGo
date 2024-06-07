package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// G -
func G(tags ...LuxeGo.Content) *GTag {
	return &GTag{ComponentSvgTag: &ComponentSvgTag{Name: "g", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type GTag struct {
	*ComponentSvgTag
}
