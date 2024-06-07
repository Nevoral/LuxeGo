package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Mark -
func Mark(tags ...interface{}) *MarkTag {
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
	return &MarkTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "mark", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MarkTag struct {
	*ComponentHtmlTag
}
