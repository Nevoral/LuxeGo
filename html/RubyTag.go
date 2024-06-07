package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Ruby -
func Ruby(tags ...interface{}) *RubyTag {
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
	return &RubyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ruby", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type RubyTag struct {
	*ComponentHtmlTag
}
