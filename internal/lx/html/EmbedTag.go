package html

import (
	"LuxeGo/internal/lx"
)

//Embed - 
func Embed() *EmbedTag {
	return &EmbedTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "embed", Attributes: &lx.Attributes{}, Children: nil}}
}

type EmbedTag struct {
	*ComponentHtmlTag
}

//Height - 
func (e *EmbedTag) Height(value string) *EmbedTag {
	e.AddAttribute("height", value)
	return e
}

//Src - Specifies the URL of an image.
func (e *EmbedTag) Src(value string) *EmbedTag {
	e.AddAttribute("src", value)
	return e
}

//Type - Specifies the type of an <input> element.
func (e *EmbedTag) Type(value string) *EmbedTag {
	e.AddAttribute("type", value)
	return e
}

//Width - 
func (e *EmbedTag) Width(value string) *EmbedTag {
	e.AddAttribute("width", value)
	return e
}
