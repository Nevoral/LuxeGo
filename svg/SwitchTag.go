package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Switch -
func Switch(tags ...interface{}) *SwitchTag {
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
	return &SwitchTag{ComponentSvgTag: &ComponentSvgTag{Name: "switch", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SwitchTag struct {
	*ComponentSvgTag
}
