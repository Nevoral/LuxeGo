package html

import (
	"LuxeGo/internal/lx"
)

//Input - 
func Input() *InputTag {
	return &InputTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "input", Attributes: &lx.Attributes{}, Children: nil}}}
}

type InputTag struct {
	*ComponentTag
}

//Accept - Defines the file types the file input should accept
func (i *InputTag) Accept(value string) *InputTag {
	i.AddAttribute("accept", value)
	return i
}

//Alt -  Provides alternative text for an image.
func (i *InputTag) Alt(value string) *InputTag {
	i.AddAttribute("alt", value)
	return i
}

//Autocapitalize - Controls whether and how text input is automatically capitalized.
func (i *InputTag) Autocapitalize(value string) *InputTag {
	i.AddAttribute("autocapitalize", value)
	return i
}

//Autocomplete - Specifies whether a form or an input field should have autocomplete enabled.
func (i *InputTag) Autocomplete(value string) *InputTag {
	i.AddAttribute("autocomplete", value)
	return i
}

//Autofocus - Specifies that a form element should automatically get focus when the page loads.
func (i *InputTag) Autofocus(value string) *InputTag {
	i.AddAttribute("autofocus", value)
	return i
}

//Capture - 
func (i *InputTag) Capture(value string) *InputTag {
	i.AddAttribute("capture", value)
	return i
}

//Checked - Indicates whether an input element is selected by default.
func (i *InputTag) Checked() *InputTag {
	i.AddAttribute("checked", "")
	return i
}

//Dirname - 
func (i *InputTag) Dirname(value string) *InputTag {
	i.AddAttribute("dirname", value)
	return i
}

//Disabled - Indicates that the user cannot interact with the element.
func (i *InputTag) Disabled() *InputTag {
	i.AddAttribute("disabled", "")
	return i
}

//Form - 
func (i *InputTag) Form(value string) *InputTag {
	i.AddAttribute("form", value)
	return i
}

//Formaction - 
func (i *InputTag) Formaction(value string) *InputTag {
	i.AddAttribute("formaction", value)
	return i
}

//Formenctype - 
func (i *InputTag) Formenctype(value string) *InputTag {
	i.AddAttribute("formenctype", value)
	return i
}

//Formmethod - 
func (i *InputTag) Formmethod(value string) *InputTag {
	i.AddAttribute("formmethod", value)
	return i
}

//Formnovalidate - 
func (i *InputTag) Formnovalidate(value string) *InputTag {
	i.AddAttribute("formnovalidate", value)
	return i
}

//Formtarget - 
func (i *InputTag) Formtarget(value string) *InputTag {
	i.AddAttribute("formtarget", value)
	return i
}

//Height - 
func (i *InputTag) Height(value string) *InputTag {
	i.AddAttribute("height", value)
	return i
}

//Inputmode - 
func (i *InputTag) Inputmode(value string) *InputTag {
	i.AddAttribute("inputmode", value)
	return i
}

//List - 
func (i *InputTag) List(value string) *InputTag {
	i.AddAttribute("list", value)
	return i
}

//Max - 
func (i *InputTag) Max(value string) *InputTag {
	i.AddAttribute("max", value)
	return i
}

//Maxlength - Specify the maximum lengths of text in an input field.
func (i *InputTag) Maxlength(value string) *InputTag {
	i.AddAttribute("maxlength", value)
	return i
}

//Min - 
func (i *InputTag) Min(value string) *InputTag {
	i.AddAttribute("min", value)
	return i
}

//Minlength - Specify the minimum lengths of text in an input field.
func (i *InputTag) Minlength(value string) *InputTag {
	i.AddAttribute("minlength", value)
	return i
}

//Multiple - Indicates that multiple values can be entered or selected.
func (i *InputTag) Multiple() *InputTag {
	i.AddAttribute("multiple", "")
	return i
}

//Name - 
func (i *InputTag) Name(value string) *InputTag {
	i.AddAttribute("name", value)
	return i
}

//Pattern - 
func (i *InputTag) Pattern(value string) *InputTag {
	i.AddAttribute("pattern", value)
	return i
}

//Placeholder - Provides a hint to the user about what can be entered in the field.
func (i *InputTag) Placeholder(value string) *InputTag {
	i.AddAttribute("placeholder", value)
	return i
}

//Popovertarget - 
func (i *InputTag) Popovertarget(value string) *InputTag {
	i.AddAttribute("popovertarget", value)
	return i
}

//Popovertargetaction - 
func (i *InputTag) Popovertargetaction(value string) *InputTag {
	i.AddAttribute("popovertargetaction", value)
	return i
}

//Readonly - Indicates that the input field is read-only.
func (i *InputTag) Readonly() *InputTag {
	i.AddAttribute("readonly", "")
	return i
}

//Required - Specifies that an input field must be filled out before submitting the form.
func (i *InputTag) Required() *InputTag {
	i.AddAttribute("required", "")
	return i
}

//Size - 
func (i *InputTag) Size(value string) *InputTag {
	i.AddAttribute("size", value)
	return i
}

//Src - Specifies the URL of an image.
func (i *InputTag) Src(value string) *InputTag {
	i.AddAttribute("src", value)
	return i
}

//Step - 
func (i *InputTag) Step(value string) *InputTag {
	i.AddAttribute("step", value)
	return i
}

//Tabindex - 
func (i *InputTag) Tabindex(value string) *InputTag {
	i.AddAttribute("tabindex", value)
	return i
}

//Title - 
func (i *InputTag) Title(value string) *InputTag {
	i.AddAttribute("title", value)
	return i
}

//Type - Specifies the type of an <input> element.
func (i *InputTag) Type(value string) *InputTag {
	i.AddAttribute("type", value)
	return i
}

//Value - Specifies the value of an <input> element.
func (i *InputTag) Value(value string) *InputTag {
	i.AddAttribute("value", value)
	return i
}

//Width - 
func (i *InputTag) Width(value string) *InputTag {
	i.AddAttribute("width", value)
	return i
}

//Autocorrect - 
func (i *InputTag) Autocorrect(value string) *InputTag {
	i.AddAttribute("autocorrect", value)
	return i
}

//Incremental - 
func (i *InputTag) Incremental(value string) *InputTag {
	i.AddAttribute("incremental", value)
	return i
}

//Orient - 
func (i *InputTag) Orient(value string) *InputTag {
	i.AddAttribute("orient", value)
	return i
}

//Results - 
func (i *InputTag) Results(value string) *InputTag {
	i.AddAttribute("results", value)
	return i
}

//Webkitdirectory - 
func (i *InputTag) Webkitdirectory(value string) *InputTag {
	i.AddAttribute("webkitdirectory", value)
	return i
}
