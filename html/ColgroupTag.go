package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Colgroup -
func Colgroup(tags ...interface{}) *ColgroupTag {
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
	return &ColgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "colgroup", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ColgroupTag struct {
	*ComponentHtmlTag
}

// Span -
func (c *ColgroupTag) Span(value string) *ColgroupTag {
	c.AddAttribute("span", value)
	return c
}
