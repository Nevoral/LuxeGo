package html

import (
	"LuxeGo/internal/lx"
)

// Table -
func Table(tags ...lx.Content) *TableTag {
	return &TableTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "table", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TableTag struct {
	*ComponentHtmlTag
}
