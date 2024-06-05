package html

import (
	"LuxeGo/internal/lx"
)

//Label - 
func Label(tags ...lx.Content) *LabelTag {
	return &LabelTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "label", Attributes: &lx.Attributes{}, Children: &tags}}
}

type LabelTag struct {
	*ComponentHtmlTag
}

//For - 
func (l *LabelTag) For(value string) *LabelTag {
	l.AddAttribute("for", value)
	return l
}
