package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// U -
func U(tags ...interface{}) *UTag {
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
	return &UTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "u", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type UTag struct {
	*ComponentHtmlTag
}
