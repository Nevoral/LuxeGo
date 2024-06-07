package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Bdi -
func Bdi(tags ...interface{}) *BdiTag {
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
	return &BdiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "bdi", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type BdiTag struct {
	*ComponentHtmlTag
}
