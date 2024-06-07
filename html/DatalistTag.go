package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Datalist -
func Datalist(tags ...interface{}) *DatalistTag {
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
	return &DatalistTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "datalist", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DatalistTag struct {
	*ComponentHtmlTag
}
