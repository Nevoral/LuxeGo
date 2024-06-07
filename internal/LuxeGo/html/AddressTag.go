package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Address -
func Address(tags ...LuxeGo.Content) *AddressTag {
	return &AddressTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "address", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type AddressTag struct {
	*ComponentHtmlTag
}
