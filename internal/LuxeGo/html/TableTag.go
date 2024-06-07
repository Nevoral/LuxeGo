package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Table -
func Table(tags ...LuxeGo.Content) *TableTag {
	return &TableTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "table", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TableTag struct {
	*ComponentHtmlTag
}
