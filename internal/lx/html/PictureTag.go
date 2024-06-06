package html

import (
	"LuxeGo/internal/lx"
)

// Picture -
func Picture(tags ...lx.Content) *PictureTag {
	return &PictureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "picture", Attributes: &lx.Attributes{}, Children: &tags}}
}

type PictureTag struct {
	*ComponentHtmlTag
}
