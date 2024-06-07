package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Picture -
func Picture(tags ...LuxeGo.Content) *PictureTag {
	return &PictureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "picture", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type PictureTag struct {
	*ComponentHtmlTag
}
