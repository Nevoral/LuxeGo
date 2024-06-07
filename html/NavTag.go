package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Nav -
func Nav(tags ...LuxeGo.Content) *NavTag {
	return &NavTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "nav", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type NavTag struct {
	*ComponentHtmlTag
}
