package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// FeMerge -
func FeMerge(tags ...LuxeGo.Content) *FeMergeTag {
	return &FeMergeTag{ComponentSvgTag: &ComponentSvgTag{Name: "femerge", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type FeMergeTag struct {
	*ComponentSvgTag
}
