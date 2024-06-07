package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Em -
func Em(tags ...interface{}) *EmTag {
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
	return &EmTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "em", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type EmTag struct {
	*ComponentHtmlTag
}
