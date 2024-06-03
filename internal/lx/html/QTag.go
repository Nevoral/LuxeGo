package html

import (
	"LuxeGo/internal/lx"
)

//Q - 
func Q(tags ...lx.Content) *QTag {
	return &QTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "q", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type QTag struct {
	*ComponentTag
}

//Cite - Specifies a reference to the source of a quote or information.
func (q *QTag) Cite(value string) *QTag {
	q.AddAttribute("cite", value)
	return q
}
