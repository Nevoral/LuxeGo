package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Figure -
func Figure(tags ...interface{}) *FigureTag {
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
	return &FigureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figure", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type FigureTag struct {
	*ComponentHtmlTag
}
