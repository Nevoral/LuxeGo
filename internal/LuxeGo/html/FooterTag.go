package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Footer -
func Footer(tags ...LuxeGo.Content) *FooterTag {
	return &FooterTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "footer", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type FooterTag struct {
	*ComponentHtmlTag
}
