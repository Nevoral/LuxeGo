package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Fieldset -
func Fieldset(tags ...interface{}) *FieldsetTag {
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
	return &FieldsetTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fieldset", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type FieldsetTag struct {
	*ComponentHtmlTag
}

// Disabled - Indicates that the user cannot interact with the element.
func (f *FieldsetTag) Disabled() *FieldsetTag {
	f.AddAttribute("disabled", "")
	return f
}

// Form -
func (f *FieldsetTag) Form(value string) *FieldsetTag {
	f.AddAttribute("form", value)
	return f
}

// Name -
func (f *FieldsetTag) Name(value string) *FieldsetTag {
	f.AddAttribute("name", value)
	return f
}
