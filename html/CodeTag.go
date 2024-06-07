package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Code -
func Code(tags ...interface{}) *CodeTag {
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
	return &CodeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "code", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type CodeTag struct {
	*ComponentHtmlTag
}
