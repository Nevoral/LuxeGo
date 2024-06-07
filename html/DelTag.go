package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Del -
func Del(tags ...interface{}) *DelTag {
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
	return &DelTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "del", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DelTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (d *DelTag) Cite(value string) *DelTag {
	d.AddAttribute("cite", value)
	return d
}

// Datetime - Specifies the date and time.
func (d *DelTag) Datetime(value string) *DelTag {
	d.AddAttribute("datetime", value)
	return d
}
