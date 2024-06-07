package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeMergeNode -
func FeMergeNode() *FeMergeNodeTag {
	return &FeMergeNodeTag{ComponentSvgTag: &ComponentSvgTag{Name: "femergenode", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeMergeNodeTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeMergeNodeTag) In(value string) *FeMergeNodeTag {
	f.AddAttribute("in", value)
	return f
}
