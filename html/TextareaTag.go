package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Textarea -
func Textarea(tags ...interface{}) *TextareaTag {
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
	return &TextareaTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "textarea", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TextareaTag struct {
	*ComponentHtmlTag
}

// Autocapitalize - Controls whether and how text input is automatically capitalized.
func (t *TextareaTag) Autocapitalize(value string) *TextareaTag {
	t.AddAttribute("autocapitalize", value)
	return t
}

// Autocomplete - Specifies whether a form or an input field should have autocomplete enabled.
func (t *TextareaTag) Autocomplete(value string) *TextareaTag {
	t.AddAttribute("autocomplete", value)
	return t
}

// Autocorrect -
func (t *TextareaTag) Autocorrect(value string) *TextareaTag {
	t.AddAttribute("autocorrect", value)
	return t
}

// Autofocus - Specifies that a form element should automatically get focus when the page loads.
func (t *TextareaTag) Autofocus(value string) *TextareaTag {
	t.AddAttribute("autofocus", value)
	return t
}

// Cols - Specify the size of a textarea.
func (t *TextareaTag) Cols(value string) *TextareaTag {
	t.AddAttribute("cols", value)
	return t
}

// Dirname -
func (t *TextareaTag) Dirname(value string) *TextareaTag {
	t.AddAttribute("dirname", value)
	return t
}

// Disabled - Indicates that the user cannot interact with the element.
func (t *TextareaTag) Disabled() *TextareaTag {
	t.AddAttribute("disabled", "")
	return t
}

// Form -
func (t *TextareaTag) Form(value string) *TextareaTag {
	t.AddAttribute("form", value)
	return t
}

// Maxlength - Specify the maximum lengths of text in an input field.
func (t *TextareaTag) Maxlength(value string) *TextareaTag {
	t.AddAttribute("maxlength", value)
	return t
}

// Minlength - Specify the minimum lengths of text in an input field.
func (t *TextareaTag) Minlength(value string) *TextareaTag {
	t.AddAttribute("minlength", value)
	return t
}

// Name -
func (t *TextareaTag) Name(value string) *TextareaTag {
	t.AddAttribute("name", value)
	return t
}

// Placeholder - Provides a hint to the user about what can be entered in the field.
func (t *TextareaTag) Placeholder(value string) *TextareaTag {
	t.AddAttribute("placeholder", value)
	return t
}

// Readonly - Indicates that the input field is read-only.
func (t *TextareaTag) Readonly() *TextareaTag {
	t.AddAttribute("readonly", "")
	return t
}

// Required - Specifies that an input field must be filled out before submitting the form.
func (t *TextareaTag) Required() *TextareaTag {
	t.AddAttribute("required", "")
	return t
}

// Rows - Specify the size of a textarea.
func (t *TextareaTag) Rows(value string) *TextareaTag {
	t.AddAttribute("rows", value)
	return t
}

// Spellcheck -
func (t *TextareaTag) Spellcheck(value string) *TextareaTag {
	t.AddAttribute("spellcheck", value)
	return t
}

// Wrap -
func (t *TextareaTag) Wrap(value string) *TextareaTag {
	t.AddAttribute("wrap", value)
	return t
}
