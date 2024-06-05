package svg

import (
	"LuxeGo/internal/lx"
)

//Switch - 
func Switch(tags ...lx.Content) *SwitchTag {
	return &SwitchTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "switch", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SwitchTag struct {
	*ComponentHtmlTag
}
