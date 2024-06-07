package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// S -
func S(tags ...interface{}) *STag {
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
	return &STag{ComponentHtmlTag: &ComponentHtmlTag{Name: "s", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type STag struct {
	*ComponentHtmlTag
}
