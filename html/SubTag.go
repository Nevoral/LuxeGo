package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Sub -
func Sub(tags ...interface{}) *SubTag {
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
	return &SubTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sub", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SubTag struct {
	*ComponentHtmlTag
}
