package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Dl -
func Dl(tags ...LuxeGo.Content) *DlTag {
	return &DlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dl", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DlTag struct {
	*ComponentHtmlTag
}
