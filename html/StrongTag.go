package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Strong -
func Strong(tags ...interface{}) *StrongTag {
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
	return &StrongTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "strong", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type StrongTag struct {
	*ComponentHtmlTag
}
