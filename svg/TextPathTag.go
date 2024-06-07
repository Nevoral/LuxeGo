package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// TextPath -
func TextPath() *TextPathTag {
	return &TextPathTag{ComponentSvgTag: &ComponentSvgTag{Name: "textpath", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type TextPathTag struct {
	*ComponentSvgTag
}

// Href -
func (t *TextPathTag) Href(value string) *TextPathTag {
	t.AddAttribute("href", value)
	return t
}

// StartOffset -
func (t *TextPathTag) StartOffset(value string) *TextPathTag {
	t.AddAttribute("startoffset", value)
	return t
}

// Method -
func (t *TextPathTag) Method(value string) *TextPathTag {
	t.AddAttribute("method", value)
	return t
}

// Spacing -
func (t *TextPathTag) Spacing(value string) *TextPathTag {
	t.AddAttribute("spacing", value)
	return t
}
