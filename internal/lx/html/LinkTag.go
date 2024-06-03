package html

import (
	"LuxeGo/internal/lx"
)

//Link - 
func Link() *LinkTag {
	return &LinkTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "link", Attributes: &lx.Attributes{}, Children: nil}}}
}

type LinkTag struct {
	*ComponentTag
}

//As - 
func (l *LinkTag) As(value string) *LinkTag {
	l.AddAttribute("as", value)
	return l
}

//Blocking - 
func (l *LinkTag) Blocking(value string) *LinkTag {
	l.AddAttribute("blocking", value)
	return l
}

//Crossorigin - How to handle cross-origin requests for the element.
func (l *LinkTag) Crossorigin(value string) *LinkTag {
	l.AddAttribute("crossorigin", value)
	return l
}

//Disabled - Indicates that the user cannot interact with the element.
func (l *LinkTag) Disabled() *LinkTag {
	l.AddAttribute("disabled", "")
	return l
}

//Fetchpriority - 
func (l *LinkTag) Fetchpriority(value string) *LinkTag {
	l.AddAttribute("fetchpriority", value)
	return l
}

//Href - Specifies the URL of the page the link goes to.
func (l *LinkTag) Href(value string) *LinkTag {
	l.AddAttribute("href", value)
	return l
}

//Hreflang - 
func (l *LinkTag) Hreflang(value string) *LinkTag {
	l.AddAttribute("hreflang", value)
	return l
}

//Imagesizes - 
func (l *LinkTag) Imagesizes(value string) *LinkTag {
	l.AddAttribute("imagesizes", value)
	return l
}

//Imagesrcset - 
func (l *LinkTag) Imagesrcset(value string) *LinkTag {
	l.AddAttribute("imagesrcset", value)
	return l
}

//Integrity - 
func (l *LinkTag) Integrity(value string) *LinkTag {
	l.AddAttribute("integrity", value)
	return l
}

//Media - 
func (l *LinkTag) Media(value string) *LinkTag {
	l.AddAttribute("media", value)
	return l
}

//Referrerpolicy - 
func (l *LinkTag) Referrerpolicy(value string) *LinkTag {
	l.AddAttribute("referrerpolicy", value)
	return l
}

//Rel - 
func (l *LinkTag) Rel(value string) *LinkTag {
	l.AddAttribute("rel", value)
	return l
}

//Sizes - 
func (l *LinkTag) Sizes(value string) *LinkTag {
	l.AddAttribute("sizes", value)
	return l
}

//Title - 
func (l *LinkTag) Title(value string) *LinkTag {
	l.AddAttribute("title", value)
	return l
}

//Type - Specifies the type of an <input> element.
func (l *LinkTag) Type(value string) *LinkTag {
	l.AddAttribute("type", value)
	return l
}
