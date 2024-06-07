package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// View -
func View(tags ...interface{}) *ViewTag {
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
	return &ViewTag{ComponentSvgTag: &ComponentSvgTag{Name: "view", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ViewTag struct {
	*ComponentSvgTag
}

// ViewBox -
func (v *ViewTag) ViewBox(value string) *ViewTag {
	v.AddAttribute("viewbox", value)
	return v
}

// PreserveAspectRatio -
func (v *ViewTag) PreserveAspectRatio(value string) *ViewTag {
	v.AddAttribute("preserveaspectratio", value)
	return v
}
