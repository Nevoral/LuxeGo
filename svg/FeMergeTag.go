package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// FeMerge -
func FeMerge(tags ...interface{}) *FeMergeTag {
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
	return &FeMergeTag{ComponentSvgTag: &ComponentSvgTag{Name: "femerge", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type FeMergeTag struct {
	*ComponentSvgTag
}
