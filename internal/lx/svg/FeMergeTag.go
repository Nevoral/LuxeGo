package svg

import (
	"LuxeGo/internal/lx"
)

//FeMerge - 
func FeMerge(tags ...lx.Content) *FeMergeTag {
	return &FeMergeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "femerge", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FeMergeTag struct {
	*ComponentHtmlTag
}
