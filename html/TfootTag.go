package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Tfoot -
func Tfoot(tags ...interface{}) *TfootTag {
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
	return &TfootTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tfoot", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TfootTag struct {
	*ComponentHtmlTag
}
