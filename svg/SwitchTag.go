package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Switch -
func Switch(tags ...LuxeGo.Content) *SwitchTag {
	return &SwitchTag{ComponentSvgTag: &ComponentSvgTag{Name: "switch", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SwitchTag struct {
	*ComponentSvgTag
}
