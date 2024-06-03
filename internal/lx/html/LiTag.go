package html

import (
	"LuxeGo/internal/lx"
)

//Li - 
func Li(tags ...lx.Content) *LiTag {
	return &LiTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "li", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type LiTag struct {
	*ComponentTag
}

//Value - Specifies the value of an <input> element.
func (l *LiTag) Value(value string) *LiTag {
	l.AddAttribute("value", value)
	return l
}
