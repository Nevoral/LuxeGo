package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// ForeignObject -
func ForeignObject(tags ...interface{}) *ForeignObjectTag {
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
	return &ForeignObjectTag{ComponentSvgTag: &ComponentSvgTag{Name: "foreignobject", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ForeignObjectTag struct {
	*ComponentSvgTag
}

// X -
func (f *ForeignObjectTag) X(value string) *ForeignObjectTag {
	f.AddAttribute("x", value)
	return f
}

// Y -
func (f *ForeignObjectTag) Y(value string) *ForeignObjectTag {
	f.AddAttribute("y", value)
	return f
}

// Width -
func (f *ForeignObjectTag) Width(value string) *ForeignObjectTag {
	f.AddAttribute("width", value)
	return f
}

// Height -
func (f *ForeignObjectTag) Height(value string) *ForeignObjectTag {
	f.AddAttribute("height", value)
	return f
}
