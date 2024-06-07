package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Label -
func Label(tags ...LuxeGo.Content) *LabelTag {
	return &LabelTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "label", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type LabelTag struct {
	*ComponentHtmlTag
}

// For -
func (l *LabelTag) For(value string) *LabelTag {
	l.AddAttribute("for", value)
	return l
}
