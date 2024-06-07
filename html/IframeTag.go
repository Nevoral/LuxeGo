package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Iframe -
func Iframe(tags ...interface{}) *IframeTag {
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
	return &IframeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "iframe", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type IframeTag struct {
	*ComponentHtmlTag
}

// Allow -
func (i *IframeTag) Allow(value string) *IframeTag {
	i.AddAttribute("allow", value)
	return i
}

// Allowfullscreen - Allows a frame to be displayed in full-screen mode.
func (i *IframeTag) Allowfullscreen(value string) *IframeTag {
	i.AddAttribute("allowfullscreen", value)
	return i
}

// Allowpaymentrequest -
func (i *IframeTag) Allowpaymentrequest(value string) *IframeTag {
	i.AddAttribute("allowpaymentrequest", value)
	return i
}

// Credentialless -
func (i *IframeTag) Credentialless(value string) *IframeTag {
	i.AddAttribute("credentialless", value)
	return i
}

// Csp -
func (i *IframeTag) Csp(value string) *IframeTag {
	i.AddAttribute("csp", value)
	return i
}

// Height -
func (i *IframeTag) Height(value string) *IframeTag {
	i.AddAttribute("height", value)
	return i
}

// Loading -
func (i *IframeTag) Loading(value string) *IframeTag {
	i.AddAttribute("loading", value)
	return i
}

// Name -
func (i *IframeTag) Name(value string) *IframeTag {
	i.AddAttribute("name", value)
	return i
}

// Refferrerpolicy -
func (i *IframeTag) Refferrerpolicy(value string) *IframeTag {
	i.AddAttribute("refferrerpolicy", value)
	return i
}

// Sandbox -
func (i *IframeTag) Sandbox(value string) *IframeTag {
	i.AddAttribute("sandbox", value)
	return i
}

// Src - Specifies the URL of an image.
func (i *IframeTag) Src(value string) *IframeTag {
	i.AddAttribute("src", value)
	return i
}

// Srcdoc -
func (i *IframeTag) Srcdoc(value string) *IframeTag {
	i.AddAttribute("srcdoc", value)
	return i
}

// Width -
func (i *IframeTag) Width(value string) *IframeTag {
	i.AddAttribute("width", value)
	return i
}
