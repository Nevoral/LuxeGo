package html

import (
	"LuxeGo/internal/lx"
)

// Q -
func Q(tags ...lx.Content) *QTag {
	return &QTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "q", Attributes: &lx.Attributes{}, Children: &tags}}
}

type QTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (q *QTag) Cite(value string) *QTag {
	q.AddAttribute("cite", value)
	return q
}
