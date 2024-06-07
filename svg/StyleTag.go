package svg

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
	return &StyleTag{ComponentSvgTag: &ComponentSvgTag{Name: "style", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type StyleTag struct {
	*ComponentSvgTag
}

// Type -
func (s *StyleTag) Type(value string) *StyleTag {
	s.AddAttribute("type", value)
	return s
}

// Media -
func (s *StyleTag) Media(value string) *StyleTag {
	s.AddAttribute("media", value)
	return s
}

// Title -
func (s *StyleTag) Title(value string) *StyleTag {
	s.AddAttribute("title", value)
	return s
}
