package html

import (
	"LuxeGo/internal/lx"
)

// Address -
func Address(tags ...lx.Content) *AddressTag {
	return &AddressTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "address", Attributes: &lx.Attributes{}, Children: &tags}}
}

type AddressTag struct {
	*ComponentHtmlTag
}
