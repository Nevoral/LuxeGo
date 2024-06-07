package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Label -
func Label(tags ...interface{}) *LabelTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &LabelTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "label", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type LabelTag struct {
	*ComponentHtmlTag
}

// For -
func (l *LabelTag) For(value string) *LabelTag {
	l.AddAttribute("for", value)
	return l
}
