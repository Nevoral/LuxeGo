package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Form -
func Form(tags ...interface{}) *FormTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &FormTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "form", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type FormTag struct {
	*ComponentHtmlTag
}

// Accept - Defines the file types the file input should accept
func (f *FormTag) Accept(value string) *FormTag {
	f.AddAttribute("accept", value)
	return f
}

// AcceptCharset - Specifies the character encodings that are to be used for the form submission
func (f *FormTag) AcceptCharset(value string) *FormTag {
	f.AddAttribute("accept-charset", value)
	return f
}

// Autocapitalize - Controls whether and how text input is automatically capitalized.
func (f *FormTag) Autocapitalize(value string) *FormTag {
	f.AddAttribute("autocapitalize", value)
	return f
}

// Autocomplete - Specifies whether a form or an input field should have autocomplete enabled.
func (f *FormTag) Autocomplete(value string) *FormTag {
	f.AddAttribute("autocomplete", value)
	return f
}

// Name -
func (f *FormTag) Name(value string) *FormTag {
	f.AddAttribute("name", value)
	return f
}

// Rel -
func (f *FormTag) Rel(value string) *FormTag {
	f.AddAttribute("rel", value)
	return f
}

// Action - Specifies where to send the form-data when a form is submitted.
func (f *FormTag) Action(value string) *FormTag {
	f.AddAttribute("action", value)
	return f
}

// Enctype -
func (f *FormTag) Enctype(value string) *FormTag {
	f.AddAttribute("enctype", value)
	return f
}

// Method -
func (f *FormTag) Method(value string) *FormTag {
	f.AddAttribute("method", value)
	return f
}

// Novalidate -
func (f *FormTag) Novalidate() *FormTag {
	f.AddAttribute("novalidate", "")
	return f
}

// Target -
func (f *FormTag) Target(value string) *FormTag {
	f.AddAttribute("target", value)
	return f
}
