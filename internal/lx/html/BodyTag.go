package html

import (
	"LuxeGo/internal/lx"
)

//Body - 
func Body(tags ...lx.Content) *BodyTag {
	return &BodyTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "body", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type BodyTag struct {
	*ComponentTag
}

//Alink - Sets the color of an active link (deprecated in HTML5).
func (b *BodyTag) Alink(value string) *BodyTag {
	b.AddAttribute("alink", value)
	return b
}

//Background - 
func (b *BodyTag) Background(value string) *BodyTag {
	b.AddAttribute("background", value)
	return b
}

//Bgcolor - 
func (b *BodyTag) Bgcolor(value string) *BodyTag {
	b.AddAttribute("bgcolor", value)
	return b
}

//Bottommargin - 
func (b *BodyTag) Bottommargin(value string) *BodyTag {
	b.AddAttribute("bottommargin", value)
	return b
}

//Leftmargin - 
func (b *BodyTag) Leftmargin(value string) *BodyTag {
	b.AddAttribute("leftmargin", value)
	return b
}

//Link - 
func (b *BodyTag) Link(value string) *BodyTag {
	b.AddAttribute("link", value)
	return b
}

//Rightmargin - 
func (b *BodyTag) Rightmargin(value string) *BodyTag {
	b.AddAttribute("rightmargin", value)
	return b
}

//Text - 
func (b *BodyTag) Text(value string) *BodyTag {
	b.AddAttribute("text", value)
	return b
}

//Topmargin - 
func (b *BodyTag) Topmargin(value string) *BodyTag {
	b.AddAttribute("topmargin", value)
	return b
}

//Vlink - 
func (b *BodyTag) Vlink(value string) *BodyTag {
	b.AddAttribute("vlink", value)
	return b
}
