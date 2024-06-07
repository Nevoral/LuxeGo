package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Li -
func Li(tags ...interface{}) *LiTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &LiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "li", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type LiTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (l *LiTag) Value(value string) *LiTag {
	l.AddAttribute("value", value)
	return l
}
