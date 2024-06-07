package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H3 -
func H3(tags ...interface{}) *H3Tag {
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
	return &H3Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h3", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H3Tag struct {
	*ComponentHtmlTag
}
