package html

import (
	"LuxeGo/internal/lx"
)

//Progress - 
func Progress(tags ...lx.Content) *ProgressTag {
	return &ProgressTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "progress", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ProgressTag struct {
	*ComponentHtmlTag
}

//Max - 
func (p *ProgressTag) Max(value string) *ProgressTag {
	p.AddAttribute("max", value)
	return p
}

//Value - Specifies the value of an <input> element.
func (p *ProgressTag) Value(value string) *ProgressTag {
	p.AddAttribute("value", value)
	return p
}
