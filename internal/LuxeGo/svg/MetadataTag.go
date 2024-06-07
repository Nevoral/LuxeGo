package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Metadata -
func Metadata(tags ...LuxeGo.Content) *MetadataTag {
	return &MetadataTag{ComponentSvgTag: &ComponentSvgTag{Name: "metadata", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MetadataTag struct {
	*ComponentSvgTag
}
