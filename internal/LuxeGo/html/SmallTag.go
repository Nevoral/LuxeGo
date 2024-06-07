package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Small -
func Small(tags ...LuxeGo.Content) *SmallTag {
	return &SmallTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "small", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SmallTag struct {
	*ComponentHtmlTag
}
