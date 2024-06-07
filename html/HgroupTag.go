package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Hgroup -
func Hgroup(tags ...interface{}) *HgroupTag {
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
	return &HgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hgroup", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HgroupTag struct {
	*ComponentHtmlTag
}
