package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Caption -
func Caption(tags ...LuxeGo.Content) *CaptionTag {
	return &CaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "caption", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type CaptionTag struct {
	*ComponentHtmlTag
}
