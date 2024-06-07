package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Figcaption -
func Figcaption(tags ...interface{}) *FigcaptionTag {
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
	return &FigcaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figcaption", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type FigcaptionTag struct {
	*ComponentHtmlTag
}
