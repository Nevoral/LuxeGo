package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Q -
func Q(tags ...LuxeGo.Content) *QTag {
	return &QTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "q", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type QTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (q *QTag) Cite(value string) *QTag {
	q.AddAttribute("cite", value)
	return q
}
