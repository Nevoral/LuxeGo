package html

import (
	"LuxeGo/internal/lx"
)

//Picture - 
func Picture(tags ...lx.Content) *PictureTag {
	return &PictureTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "picture", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type PictureTag struct {
	*ComponentTag
}
