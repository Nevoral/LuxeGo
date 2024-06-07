package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Ins -
func Ins(tags ...interface{}) *InsTag {
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
	return &InsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ins", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type InsTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (i *InsTag) Cite(value string) *InsTag {
	i.AddAttribute("cite", value)
	return i
}

// Datetime - Specifies the date and time.
func (i *InsTag) Datetime(value string) *InsTag {
	i.AddAttribute("datetime", value)
	return i
}
