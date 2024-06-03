package html

import (
	"LuxeGo/internal/lx"
)

//Ruby - 
func Ruby(tags ...lx.Content) *RubyTag {
	return &RubyTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "ruby", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type RubyTag struct {
	*ComponentTag
}
