package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Blockquote -
func Blockquote(tags ...LuxeGo.Content) *BlockquoteTag {
	return &BlockquoteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "blockquote", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type BlockquoteTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (b *BlockquoteTag) Cite(value string) *BlockquoteTag {
	b.AddAttribute("cite", value)
	return b
}
