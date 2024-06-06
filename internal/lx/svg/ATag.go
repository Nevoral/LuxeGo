package svg

import (
	"LuxeGo/internal/lx"
)

// A -
func A(tags ...lx.Content) *ATag {
	return &ATag{ComponentSvgTag: &ComponentSvgTag{Name: "a", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ATag struct {
	*ComponentSvgTag
}

// Download -
func (a *ATag) Download(value string) *ATag {
	a.AddAttribute("download", value)
	return a
}

// Href -
func (a *ATag) Href(value string) *ATag {
	a.AddAttribute("href", value)
	return a
}

// Hreflang -
func (a *ATag) Hreflang(value string) *ATag {
	a.AddAttribute("hreflang", value)
	return a
}

// Ping -
func (a *ATag) Ping(value string) *ATag {
	a.AddAttribute("ping", value)
	return a
}

// Referrerpolicy -
func (a *ATag) Referrerpolicy(value string) *ATag {
	a.AddAttribute("referrerpolicy", value)
	return a
}

// Rel -
func (a *ATag) Rel(value string) *ATag {
	a.AddAttribute("rel", value)
	return a
}

// Target -
func (a *ATag) Target(value string) *ATag {
	a.AddAttribute("target", value)
	return a
}

// Type -
func (a *ATag) Type(value string) *ATag {
	a.AddAttribute("type", value)
	return a
}
