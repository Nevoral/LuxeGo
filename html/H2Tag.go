package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H2 -
func H2(tags ...interface{}) *H2Tag {
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
	return &H2Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h2", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H2Tag struct {
	*ComponentHtmlTag
}
