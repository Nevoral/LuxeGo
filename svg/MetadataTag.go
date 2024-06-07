package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Metadata -
func Metadata(tags ...interface{}) *MetadataTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &MetadataTag{ComponentSvgTag: &ComponentSvgTag{Name: "metadata", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MetadataTag struct {
	*ComponentSvgTag
}
