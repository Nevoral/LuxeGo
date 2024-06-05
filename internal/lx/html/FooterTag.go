package html

import (
	"LuxeGo/internal/lx"
)

//Footer - 
func Footer(tags ...lx.Content) *FooterTag {
	return &FooterTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "footer", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FooterTag struct {
	*ComponentHtmlTag
}
