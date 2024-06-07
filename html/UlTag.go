package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Ul -
func Ul(tags ...interface{}) *UlTag {
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
	return &UlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ul", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type UlTag struct {
	*ComponentHtmlTag
}
