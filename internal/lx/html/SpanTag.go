package html

import (
	"LuxeGo/internal/lx"
)

// Span -
func Span(tags ...lx.Content) *SpanTag {
	return &SpanTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "span", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SpanTag struct {
	*ComponentHtmlTag
}
