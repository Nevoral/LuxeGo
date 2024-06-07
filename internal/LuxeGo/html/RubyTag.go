package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Ruby -
func Ruby(tags ...LuxeGo.Content) *RubyTag {
	return &RubyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ruby", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type RubyTag struct {
	*ComponentHtmlTag
}
