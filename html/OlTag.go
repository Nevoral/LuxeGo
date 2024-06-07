package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Ol -
func Ol(tags ...interface{}) *OlTag {
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
	return &OlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ol", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type OlTag struct {
	*ComponentHtmlTag
}

// Reversed -
func (o *OlTag) Reversed() *OlTag {
	o.AddAttribute("reversed", "")
	return o
}

// Start -
func (o *OlTag) Start(value string) *OlTag {
	o.AddAttribute("start", value)
	return o
}

// Type - Specifies the type of an <input> element.
func (o *OlTag) Type(value string) *OlTag {
	o.AddAttribute("type", value)
	return o
}
