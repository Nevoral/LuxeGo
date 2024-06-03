package html

import (
	"LuxeGo/internal/lx"
)

//Label - 
func Label(tags ...lx.Content) *LabelTag {
	return &LabelTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "label", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type LabelTag struct {
	*ComponentTag
}

//For - 
func (l *LabelTag) For(value string) *LabelTag {
	l.AddAttribute("for", value)
	return l
}
