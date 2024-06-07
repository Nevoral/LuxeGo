package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Nav -
func Nav(tags ...interface{}) *NavTag {
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
	return &NavTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "nav", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type NavTag struct {
	*ComponentHtmlTag
}
