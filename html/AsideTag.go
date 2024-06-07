package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Aside -
func Aside(tags ...interface{}) *AsideTag {
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
	return &AsideTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "aside", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type AsideTag struct {
	*ComponentHtmlTag
}
