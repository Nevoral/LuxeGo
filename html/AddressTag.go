package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Address -
func Address(tags ...interface{}) *AddressTag {
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
	return &AddressTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "address", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type AddressTag struct {
	*ComponentHtmlTag
}
