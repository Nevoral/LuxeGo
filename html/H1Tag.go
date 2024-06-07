package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H1 -
func H1(tags ...interface{}) *H1Tag {
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
	return &H1Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h1", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H1Tag struct {
	*ComponentHtmlTag
}
