package html

import (
	"LuxeGo/internal/lx"
)

// Button -
func Button(tags ...lx.Content) *ButtonTag {
	return &ButtonTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "button", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ButtonTag struct {
	*ComponentHtmlTag
}

// Autofocus - Specifies that a form element should automatically get focus when the page loads.
func (b *ButtonTag) Autofocus(value string) *ButtonTag {
	b.AddAttribute("autofocus", value)
	return b
}

// Autocomplete - Specifies whether a form or an input field should have autocomplete enabled.
func (b *ButtonTag) Autocomplete(value string) *ButtonTag {
	b.AddAttribute("autocomplete", value)
	return b
}

// Disabled - Indicates that the user cannot interact with the element.
func (b *ButtonTag) Disabled() *ButtonTag {
	b.AddAttribute("disabled", "")
	return b
}

// Form -
func (b *ButtonTag) Form(value string) *ButtonTag {
	b.AddAttribute("form", value)
	return b
}

// Formaction -
func (b *ButtonTag) Formaction(value string) *ButtonTag {
	b.AddAttribute("formaction", value)
	return b
}

// Formenctype -
func (b *ButtonTag) Formenctype(value string) *ButtonTag {
	b.AddAttribute("formenctype", value)
	return b
}

// Formmethod -
func (b *ButtonTag) Formmethod(value string) *ButtonTag {
	b.AddAttribute("formmethod", value)
	return b
}

// Formnovalidate -
func (b *ButtonTag) Formnovalidate(value string) *ButtonTag {
	b.AddAttribute("formnovalidate", value)
	return b
}

// Formtarget -
func (b *ButtonTag) Formtarget(value string) *ButtonTag {
	b.AddAttribute("formtarget", value)
	return b
}

// Name -
func (b *ButtonTag) Name(value string) *ButtonTag {
	b.AddAttribute("name", value)
	return b
}

// Popovertarget -
func (b *ButtonTag) Popovertarget(value string) *ButtonTag {
	b.AddAttribute("popovertarget", value)
	return b
}

// Popovertargetaction -
func (b *ButtonTag) Popovertargetaction(value string) *ButtonTag {
	b.AddAttribute("popovertargetaction", value)
	return b
}

// Type - Specifies the type of an <input> element.
func (b *ButtonTag) Type(value string) *ButtonTag {
	b.AddAttribute("type", value)
	return b
}

// Value - Specifies the value of an <input> element.
func (b *ButtonTag) Value(value string) *ButtonTag {
	b.AddAttribute("value", value)
	return b
}
