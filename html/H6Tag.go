package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H6 -
func H6(tags ...interface{}) *H6Tag {
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
	return &H6Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h6", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H6Tag struct {
	*ComponentHtmlTag
}
