package svg

import (
	"LuxeGo/internal/lx"
)

//FeMergeNode - 
func FeMergeNode() *FeMergeNodeTag {
	return &FeMergeNodeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "femergenode", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeMergeNodeTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeMergeNodeTag) In(value string) *FeMergeNodeTag {
	f.AddAttribute("in", value)
	return f
}
