package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// H5 -
func H5(tags ...LuxeGo.Content) *H5Tag {
	return &H5Tag{ComponentHtmlTag: &ComponentHtmlTag{Name: "h5", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type H5Tag struct {
	*ComponentHtmlTag
}
