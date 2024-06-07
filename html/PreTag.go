package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Pre -
func Pre(tags ...interface{}) *PreTag {
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
	return &PreTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "pre", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type PreTag struct {
	*ComponentHtmlTag
}
