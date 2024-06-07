package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Main -
func Main(tags ...LuxeGo.Content) *MainTag {
	return &MainTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "main", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MainTag struct {
	*ComponentHtmlTag
}
