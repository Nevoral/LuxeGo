package html

import (
	"github.com/Nevoral/LuxeGo"
)

// I -
func I(tags ...LuxeGo.Content) *ITag {
	return &ITag{ComponentHtmlTag: &ComponentHtmlTag{Name: "i", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ITag struct {
	*ComponentHtmlTag
}
