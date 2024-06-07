package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Picture -
func Picture(tags ...interface{}) *PictureTag {
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
	return &PictureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "picture", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type PictureTag struct {
	*ComponentHtmlTag
}
