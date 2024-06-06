package svg

import (
	"LuxeGo/internal/lx"
)

// Switch -
func Switch(tags ...lx.Content) *SwitchTag {
	return &SwitchTag{ComponentSvgTag: &ComponentSvgTag{Name: "switch", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SwitchTag struct {
	*ComponentSvgTag
}
