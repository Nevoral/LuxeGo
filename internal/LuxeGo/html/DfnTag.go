package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Dfn -
func Dfn(tags ...LuxeGo.Content) *DfnTag {
	return &DfnTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dfn", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DfnTag struct {
	*ComponentHtmlTag
}
