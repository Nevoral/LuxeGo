package html

import (
	"LuxeGo/internal/lx"
)

// Ruby -
func Ruby(tags ...lx.Content) *RubyTag {
	return &RubyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ruby", Attributes: &lx.Attributes{}, Children: &tags}}
}

type RubyTag struct {
	*ComponentHtmlTag
}
