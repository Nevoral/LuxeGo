package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Legend -
func Legend(tags ...interface{}) *LegendTag {
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
	return &LegendTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "legend", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type LegendTag struct {
	*ComponentHtmlTag
}
