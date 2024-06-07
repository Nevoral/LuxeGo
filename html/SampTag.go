package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Samp -
func Samp(tags ...interface{}) *SampTag {
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
	return &SampTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "samp", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SampTag struct {
	*ComponentHtmlTag
}
