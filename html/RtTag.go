package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Rt -
func Rt(tags ...interface{}) *RtTag {
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
	return &RtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rt", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type RtTag struct {
	*ComponentHtmlTag
}
