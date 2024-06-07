package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Small -
func Small(tags ...interface{}) *SmallTag {
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
	return &SmallTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "small", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SmallTag struct {
	*ComponentHtmlTag
}
