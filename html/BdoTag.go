package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Bdo -
func Bdo(tags ...LuxeGo.Content) *BdoTag {
	return &BdoTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "bdo", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type BdoTag struct {
	*ComponentHtmlTag
}
