package svg

import (
	"LuxeGo/internal/lx"
)

//TextPath - 
func TextPath() *TextPathTag {
	return &TextPathTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "textpath", Attributes: &lx.Attributes{}, Children: nil}}
}

type TextPathTag struct {
	*ComponentHtmlTag
}

//Href - 
func (t *TextPathTag) Href(value string) *TextPathTag {
	t.AddAttribute("href", value)
	return t
}

//StartOffset - 
func (t *TextPathTag) StartOffset(value string) *TextPathTag {
	t.AddAttribute("startoffset", value)
	return t
}

//Method - 
func (t *TextPathTag) Method(value string) *TextPathTag {
	t.AddAttribute("method", value)
	return t
}

//Spacing - 
func (t *TextPathTag) Spacing(value string) *TextPathTag {
	t.AddAttribute("spacing", value)
	return t
}
