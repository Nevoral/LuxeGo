package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Option -
func Option(tags ...interface{}) *OptionTag {
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
	return &OptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "option", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type OptionTag struct {
	*ComponentHtmlTag
}

// Disabled - Indicates that the user cannot interact with the element.
func (o *OptionTag) Disabled() *OptionTag {
	o.AddAttribute("disabled", "")
	return o
}

// Label -
func (o *OptionTag) Label(value string) *OptionTag {
	o.AddAttribute("label", value)
	return o
}

// Selected -
func (o *OptionTag) Selected() *OptionTag {
	o.AddAttribute("selected", "")
	return o
}

// Value - Specifies the value of an <input> element.
func (o *OptionTag) Value(value string) *OptionTag {
	o.AddAttribute("value", value)
	return o
}
