package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Q -
func Q(tags ...interface{}) *QTag {
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
	return &QTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "q", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type QTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (q *QTag) Cite(value string) *QTag {
	q.AddAttribute("cite", value)
	return q
}
