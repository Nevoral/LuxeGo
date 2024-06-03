package html

import (
	"LuxeGo/internal/lx"
)

//Blockquote - 
func Blockquote(tags ...lx.Content) *BlockquoteTag {
	return &BlockquoteTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "blockquote", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type BlockquoteTag struct {
	*ComponentTag
}

//Cite - Specifies a reference to the source of a quote or information.
func (b *BlockquoteTag) Cite(value string) *BlockquoteTag {
	b.AddAttribute("cite", value)
	return b
}
