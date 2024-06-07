package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Search -
func Search(tags ...interface{}) *SearchTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &SearchTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "search", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SearchTag struct {
	*ComponentHtmlTag
}
