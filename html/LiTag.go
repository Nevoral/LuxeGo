package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Li -
func Li(tags ...LuxeGo.Content) *LiTag {
	return &LiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "li", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type LiTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (l *LiTag) Value(value string) *LiTag {
	l.AddAttribute("value", value)
	return l
}
