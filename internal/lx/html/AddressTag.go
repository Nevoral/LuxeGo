package html

import (
	"LuxeGo/internal/lx"
)

//Address - 
func Address(tags ...lx.Content) *AddressTag {
	return &AddressTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "address", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type AddressTag struct {
	*ComponentTag
}
