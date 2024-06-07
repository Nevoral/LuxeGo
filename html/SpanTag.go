package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Span -
func Span(tags ...LuxeGo.Content) *SpanTag {
	return &SpanTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "span", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SpanTag struct {
	*ComponentHtmlTag
}
