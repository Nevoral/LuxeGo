package html

import (
	"LuxeGo/internal/lx"
)

// Li -
func Li(tags ...lx.Content) *LiTag {
	return &LiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "li", Attributes: &lx.Attributes{}, Children: &tags}}
}

type LiTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (l *LiTag) Value(value string) *LiTag {
	l.AddAttribute("value", value)
	return l
}
