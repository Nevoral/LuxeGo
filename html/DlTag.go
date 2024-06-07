package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Dl -
func Dl(tags ...interface{}) *DlTag {
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
	return &DlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dl", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DlTag struct {
	*ComponentHtmlTag
}
