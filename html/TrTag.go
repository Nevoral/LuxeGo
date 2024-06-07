package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Tr -
func Tr(tags ...interface{}) *TrTag {
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
	return &TrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tr", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TrTag struct {
	*ComponentHtmlTag
}
