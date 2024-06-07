package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Thead -
func Thead(tags ...interface{}) *TheadTag {
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
	return &TheadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "thead", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TheadTag struct {
	*ComponentHtmlTag
}
