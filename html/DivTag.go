package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Div -
func Div(tags ...interface{}) *DivTag {
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
	return &DivTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "div", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DivTag struct {
	*ComponentHtmlTag
}
