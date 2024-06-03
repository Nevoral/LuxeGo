package html

import (
	"LuxeGo/internal/lx"
)

//Footer - 
func Footer(tags ...lx.Content) *FooterTag {
	return &FooterTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "footer", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type FooterTag struct {
	*ComponentTag
}
