package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Blockquote -
func Blockquote(tags ...interface{}) *BlockquoteTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &BlockquoteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "blockquote", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type BlockquoteTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (b *BlockquoteTag) Cite(value string) *BlockquoteTag {
	b.AddAttribute("cite", value)
	return b
}
