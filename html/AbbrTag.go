package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Abbr -
func Abbr(tags ...interface{}) *AbbrTag {
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
	return &AbbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "abbr", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type AbbrTag struct {
	*ComponentHtmlTag
}
