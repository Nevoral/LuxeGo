package svg

import (
	"LuxeGo/internal/lx"
)

// FeMergeNode -
func FeMergeNode() *FeMergeNodeTag {
	return &FeMergeNodeTag{ComponentSvgTag: &ComponentSvgTag{Name: "femergenode", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeMergeNodeTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeMergeNodeTag) In(value string) *FeMergeNodeTag {
	f.AddAttribute("in", value)
	return f
}
