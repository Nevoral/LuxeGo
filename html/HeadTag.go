package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Head -
func Head(tags ...interface{}) *HeadTag {
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
	return &HeadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "head", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HeadTag struct {
	*ComponentHtmlTag
}
