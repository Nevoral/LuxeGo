package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Dfn -
func Dfn(tags ...interface{}) *DfnTag {
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
	return &DfnTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dfn", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DfnTag struct {
	*ComponentHtmlTag
}
