package html

import (
	"LuxeGo/internal/lx"
)

// Caption -
func Caption(tags ...lx.Content) *CaptionTag {
	return &CaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "caption", Attributes: &lx.Attributes{}, Children: &tags}}
}

type CaptionTag struct {
	*ComponentHtmlTag
}
