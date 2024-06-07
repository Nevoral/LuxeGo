package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Noscript -
func Noscript(tags ...interface{}) *NoscriptTag {
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
	return &NoscriptTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "noscript", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type NoscriptTag struct {
	*ComponentHtmlTag
}
