package svg

import (
	"LuxeGo/internal/lx"
)

//Metadata - 
func Metadata(tags ...lx.Content) *MetadataTag {
	return &MetadataTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "metadata", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MetadataTag struct {
	*ComponentHtmlTag
}
