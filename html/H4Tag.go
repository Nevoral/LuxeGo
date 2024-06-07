package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H4 -
func H4(tags ...interface{}) *H4Tag {
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
	return &H4Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h4", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H4Tag struct {
	*ComponentHtmlTag
}
