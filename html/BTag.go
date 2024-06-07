package html

import (
	"github.com/Nevoral/LuxeGo"
)

// B -
func B(tags ...LuxeGo.Content) *BTag {
	return &BTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "b", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type BTag struct {
	*ComponentHtmlTag
}
