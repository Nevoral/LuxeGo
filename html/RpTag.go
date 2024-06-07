package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Rp -
func Rp(tags ...interface{}) *RpTag {
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
	return &RpTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rp", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type RpTag struct {
	*ComponentHtmlTag
}
