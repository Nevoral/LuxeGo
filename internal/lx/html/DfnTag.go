package html

import (
	"LuxeGo/internal/lx"
)

// Dfn -
func Dfn(tags ...lx.Content) *DfnTag {
	return &DfnTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dfn", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DfnTag struct {
	*ComponentHtmlTag
}
