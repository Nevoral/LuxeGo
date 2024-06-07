package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Caption -
func Caption(tags ...interface{}) *CaptionTag {
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
	return &CaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "caption", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type CaptionTag struct {
	*ComponentHtmlTag
}
