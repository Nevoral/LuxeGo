package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Optgroup -
func Optgroup(tags ...interface{}) *OptgroupTag {
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
	return &OptgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "optgroup", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type OptgroupTag struct {
	*ComponentHtmlTag
}

// Disabled - Indicates that the user cannot interact with the element.
func (o *OptgroupTag) Disabled() *OptgroupTag {
	o.AddAttribute("disabled", "")
	return o
}

// Label -
func (o *OptgroupTag) Label(value string) *OptgroupTag {
	o.AddAttribute("label", value)
	return o
}
