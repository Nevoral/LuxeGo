package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Portal -
func Portal(tags ...LuxeGo.Content) *PortalTag {
	return &PortalTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "portal", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type PortalTag struct {
	*ComponentHtmlTag
}

// Referrerpolicy -
func (p *PortalTag) Referrerpolicy(value string) *PortalTag {
	p.AddAttribute("referrerpolicy", value)
	return p
}

// Src - Specifies the URL of an image.
func (p *PortalTag) Src(value string) *PortalTag {
	p.AddAttribute("src", value)
	return p
}
