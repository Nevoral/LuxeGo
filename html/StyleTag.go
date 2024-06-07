package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Style -
func Style(tags ...interface{}) *StyleTag {
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
	return &StyleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "style", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type StyleTag struct {
	*ComponentHtmlTag
}

// Blocking -
func (s *StyleTag) Blocking(value string) *StyleTag {
	s.AddAttribute("blocking", value)
	return s
}

// Media -
func (s *StyleTag) Media(value string) *StyleTag {
	s.AddAttribute("media", value)
	return s
}

// Nonce -
func (s *StyleTag) Nonce(value string) *StyleTag {
	s.AddAttribute("nonce", value)
	return s
}

// Type - Specifies the type of an <input> element.
func (s *StyleTag) Type(value string) *StyleTag {
	s.AddAttribute("type", value)
	return s
}
