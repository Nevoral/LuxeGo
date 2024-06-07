package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Desc -
func Desc(tags ...LuxeGo.Content) *DescTag {
	return &DescTag{ComponentSvgTag: &ComponentSvgTag{Name: "desc", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DescTag struct {
	*ComponentSvgTag
}
