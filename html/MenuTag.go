package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Menu -
func Menu(tags ...interface{}) *MenuTag {
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
	return &MenuTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "menu", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MenuTag struct {
	*ComponentHtmlTag
}
