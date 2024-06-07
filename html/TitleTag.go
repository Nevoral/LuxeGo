package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Title -
func Title(tags ...interface{}) *TitleTag {
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
	return &TitleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "title", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TitleTag struct {
	*ComponentHtmlTag
}
