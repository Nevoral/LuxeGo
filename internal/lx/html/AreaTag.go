package html

import (
	"LuxeGo/internal/lx"
)

// Area -
func Area() *AreaTag {
	return &AreaTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "area", Attributes: &lx.Attributes{}, Children: nil}}
}

type AreaTag struct {
	*ComponentHtmlTag
}

// Alt -  Provides alternative text for an image.
func (a *AreaTag) Alt(value string) *AreaTag {
	a.AddAttribute("alt", value)
	return a
}

// Coords -
func (a *AreaTag) Coords(value string) *AreaTag {
	a.AddAttribute("coords", value)
	return a
}

// Download - Specifies that the target will be downloaded when a user clicks on the hyperlink.
func (a *AreaTag) Download(value string) *AreaTag {
	a.AddAttribute("download", value)
	return a
}

// Href - Specifies the URL of the page the link goes to.
func (a *AreaTag) Href(value string) *AreaTag {
	a.AddAttribute("href", value)
	return a
}

// Ping -
func (a *AreaTag) Ping(value string) *AreaTag {
	a.AddAttribute("ping", value)
	return a
}

// Referrerpolicy -
func (a *AreaTag) Referrerpolicy(value string) *AreaTag {
	a.AddAttribute("referrerpolicy", value)
	return a
}

// Rel -
func (a *AreaTag) Rel(value string) *AreaTag {
	a.AddAttribute("rel", value)
	return a
}

// Shape -
func (a *AreaTag) Shape(value string) *AreaTag {
	a.AddAttribute("shape", value)
	return a
}

// Target -
func (a *AreaTag) Target(value string) *AreaTag {
	a.AddAttribute("target", value)
	return a
}
