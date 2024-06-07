package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Output -
func Output(tags ...interface{}) *OutputTag {
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
	return &OutputTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "output", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type OutputTag struct {
	*ComponentHtmlTag
}

// For -
func (o *OutputTag) For(value string) *OutputTag {
	o.AddAttribute("for", value)
	return o
}

// Form -
func (o *OutputTag) Form(value string) *OutputTag {
	o.AddAttribute("form", value)
	return o
}

// Name -
func (o *OutputTag) Name(value string) *OutputTag {
	o.AddAttribute("name", value)
	return o
}
