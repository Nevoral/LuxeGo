package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Header -
func Header(tags ...interface{}) *HeaderTag {
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
	return &HeaderTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "header", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HeaderTag struct {
	*ComponentHtmlTag
}
