package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Script -
func Script(tags ...LuxeGo.Content) *ScriptTag {
	return &ScriptTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "script", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ScriptTag struct {
	*ComponentHtmlTag
}

// Async - Indicates that the script should be executed asynchronously.
func (s *ScriptTag) Async() *ScriptTag {
	s.AddAttribute("async", "")
	return s
}

// Blocking -
func (s *ScriptTag) Blocking(value string) *ScriptTag {
	s.AddAttribute("blocking", value)
	return s
}

// Crossorigin - How to handle cross-origin requests for the element.
func (s *ScriptTag) Crossorigin(value string) *ScriptTag {
	s.AddAttribute("crossorigin", value)
	return s
}

// Defer -
func (s *ScriptTag) Defer() *ScriptTag {
	s.AddAttribute("defer", "")
	return s
}

// Fetchpriority -
func (s *ScriptTag) Fetchpriority(value string) *ScriptTag {
	s.AddAttribute("fetchpriority", value)
	return s
}

// Integrity -
func (s *ScriptTag) Integrity(value string) *ScriptTag {
	s.AddAttribute("integrity", value)
	return s
}

// Nomodule -
func (s *ScriptTag) Nomodule(value string) *ScriptTag {
	s.AddAttribute("nomodule", value)
	return s
}

// Nonce -
func (s *ScriptTag) Nonce(value string) *ScriptTag {
	s.AddAttribute("nonce", value)
	return s
}

// Referrerpolicy -
func (s *ScriptTag) Referrerpolicy(value string) *ScriptTag {
	s.AddAttribute("referrerpolicy", value)
	return s
}

// Src - Specifies the URL of an image.
func (s *ScriptTag) Src(value string) *ScriptTag {
	s.AddAttribute("src", value)
	return s
}

// Type - Specifies the type of an <input> element.
func (s *ScriptTag) Type(value string) *ScriptTag {
	s.AddAttribute("type", value)
	return s
}
