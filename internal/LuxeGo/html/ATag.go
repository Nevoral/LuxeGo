package html

import (
	"LuxeGo/internal/LuxeGo"
)

// A -
func A(tags ...LuxeGo.Content) *ATag {
	return &ATag{ComponentHtmlTag: &ComponentHtmlTag{Name: "a", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ATag struct {
	*ComponentHtmlTag
}

// Download - Specifies that the target will be downloaded when a user clicks on the hyperlink.
func (a *ATag) Download(value string) *ATag {
	a.AddAttribute("download", value)
	return a
}

// Href - Specifies the URL of the page the link goes to.
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

// Type - Specifies the type of an <input> element.
func (a *ATag) Type(value string) *ATag {
	a.AddAttribute("type", value)
	return a
}
