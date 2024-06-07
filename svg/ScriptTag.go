package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Script -
func Script(tags ...interface{}) *ScriptTag {
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
	return &ScriptTag{ComponentSvgTag: &ComponentSvgTag{Name: "script", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ScriptTag struct {
	*ComponentSvgTag
}

// Href -
func (s *ScriptTag) Href(value string) *ScriptTag {
	s.AddAttribute("href", value)
	return s
}

// Type -
func (s *ScriptTag) Type(value string) *ScriptTag {
	s.AddAttribute("type", value)
	return s
}
