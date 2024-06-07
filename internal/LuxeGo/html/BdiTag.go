package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Bdi -
func Bdi(tags ...LuxeGo.Content) *BdiTag {
	return &BdiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "bdi", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type BdiTag struct {
	*ComponentHtmlTag
}
