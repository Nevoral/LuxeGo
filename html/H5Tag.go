package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// H5 -
func H5(tags ...interface{}) *H5Tag {
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
	return &H5Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h5", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type H5Tag struct {
	*ComponentHtmlTag
}
