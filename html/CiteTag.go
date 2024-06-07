package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Cite -
func Cite(tags ...interface{}) *CiteTag {
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
	return &CiteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "cite", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type CiteTag struct {
	*ComponentHtmlTag
}
