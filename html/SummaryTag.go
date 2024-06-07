package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Summary -
func Summary(tags ...interface{}) *SummaryTag {
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
	return &SummaryTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "summary", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SummaryTag struct {
	*ComponentHtmlTag
}
