package html

import (
	"LuxeGo/internal/lx"
)

//Span - 
func Span(tags ...lx.Content) *SpanTag {
	return &SpanTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "span", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SpanTag struct {
	*ComponentTag
}
