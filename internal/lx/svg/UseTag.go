package svg

import (
	"LuxeGo/internal/lx"
)

//Use - 
func Use() *UseTag {
	return &UseTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "use", Attributes: &lx.Attributes{}, Children: nil}}
}

type UseTag struct {
	*ComponentHtmlTag
}

//Href - 
func (u *UseTag) Href(value string) *UseTag {
	u.AddAttribute("href", value)
	return u
}

//X - 
func (u *UseTag) X(value string) *UseTag {
	u.AddAttribute("x", value)
	return u
}

//Y - 
func (u *UseTag) Y(value string) *UseTag {
	u.AddAttribute("y", value)
	return u
}

//Width - 
func (u *UseTag) Width(value string) *UseTag {
	u.AddAttribute("width", value)
	return u
}

//Height - 
func (u *UseTag) Height(value string) *UseTag {
	u.AddAttribute("height", value)
	return u
}
