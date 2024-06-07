package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Dd -
func Dd(tags ...interface{}) *DdTag {
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
	return &DdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dd", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DdTag struct {
	*ComponentHtmlTag
}
