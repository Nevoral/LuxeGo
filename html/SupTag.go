package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Sup -
func Sup(tags ...interface{}) *SupTag {
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
	return &SupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sup", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SupTag struct {
	*ComponentHtmlTag
}
