package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Body -
func Body(tags ...interface{}) *BodyTag {
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
	return &BodyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "body", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type BodyTag struct {
	*ComponentHtmlTag
}

// Alink - Sets the color of an active link (deprecated in HTML5).
func (b *BodyTag) Alink(value string) *BodyTag {
	b.AddAttribute("alink", value)
	return b
}

// Background -
func (b *BodyTag) Background(value string) *BodyTag {
	b.AddAttribute("background", value)
	return b
}

// Bgcolor -
func (b *BodyTag) Bgcolor(value string) *BodyTag {
	b.AddAttribute("bgcolor", value)
	return b
}

// Bottommargin -
func (b *BodyTag) Bottommargin(value string) *BodyTag {
	b.AddAttribute("bottommargin", value)
	return b
}

// Leftmargin -
func (b *BodyTag) Leftmargin(value string) *BodyTag {
	b.AddAttribute("leftmargin", value)
	return b
}

// Link -
func (b *BodyTag) Link(value string) *BodyTag {
	b.AddAttribute("link", value)
	return b
}

// Rightmargin -
func (b *BodyTag) Rightmargin(value string) *BodyTag {
	b.AddAttribute("rightmargin", value)
	return b
}

// Text -
func (b *BodyTag) Text(value string) *BodyTag {
	b.AddAttribute("text", value)
	return b
}

// Topmargin -
func (b *BodyTag) Topmargin(value string) *BodyTag {
	b.AddAttribute("topmargin", value)
	return b
}

// Vlink -
func (b *BodyTag) Vlink(value string) *BodyTag {
	b.AddAttribute("vlink", value)
	return b
}
