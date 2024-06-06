package html

import (
	"LuxeGo/internal/lx"
)

// Blockquote -
func Blockquote(tags ...lx.Content) *BlockquoteTag {
	return &BlockquoteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "blockquote", Attributes: &lx.Attributes{}, Children: &tags}}
}

type BlockquoteTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (b *BlockquoteTag) Cite(value string) *BlockquoteTag {
	b.AddAttribute("cite", value)
	return b
}
