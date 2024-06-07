package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Defs -
func Defs(tags ...interface{}) *DefsTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &DefsTag{ComponentSvgTag: &ComponentSvgTag{Name: "defs", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DefsTag struct {
	*ComponentSvgTag
}
