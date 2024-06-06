package svg

import (
	"LuxeGo/internal/lx"
)

// FeMerge -
func FeMerge(tags ...lx.Content) *FeMergeTag {
	return &FeMergeTag{ComponentSvgTag: &ComponentSvgTag{Name: "femerge", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FeMergeTag struct {
	*ComponentSvgTag
}
