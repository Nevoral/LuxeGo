package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Table -
func Table(tags ...interface{}) *TableTag {
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
	return &TableTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "table", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TableTag struct {
	*ComponentHtmlTag
}
