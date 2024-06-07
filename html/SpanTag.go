package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Span -
func Span(tags ...interface{}) *SpanTag {
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
	return &SpanTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "span", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SpanTag struct {
	*ComponentHtmlTag
}
