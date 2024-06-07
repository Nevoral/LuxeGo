package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Strong -
func Strong(tags ...LuxeGo.Content) *StrongTag {
	return &StrongTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "strong", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type StrongTag struct {
	*ComponentHtmlTag
}