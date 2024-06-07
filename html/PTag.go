package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// P -
func P(tags ...interface{}) *PTag {
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
	return &PTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "p", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type PTag struct {
	*ComponentHtmlTag
}
