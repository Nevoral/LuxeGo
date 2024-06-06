package html

import (
	"LuxeGo/internal/lx"
)

// Main -
func Main(tags ...lx.Content) *MainTag {
	return &MainTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "main", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MainTag struct {
	*ComponentHtmlTag
}
