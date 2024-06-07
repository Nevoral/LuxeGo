package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Symbol -
func Symbol(tags ...interface{}) *SymbolTag {
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
	return &SymbolTag{ComponentSvgTag: &ComponentSvgTag{Name: "symbol", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type SymbolTag struct {
	*ComponentSvgTag
}

// ViewBox -
func (s *SymbolTag) ViewBox(value string) *SymbolTag {
	s.AddAttribute("viewbox", value)
	return s
}

// PreserveAspectRatio -
func (s *SymbolTag) PreserveAspectRatio(value string) *SymbolTag {
	s.AddAttribute("preserveaspectratio", value)
	return s
}

// RefX -
func (s *SymbolTag) RefX(value string) *SymbolTag {
	s.AddAttribute("refx", value)
	return s
}

// RefY -
func (s *SymbolTag) RefY(value string) *SymbolTag {
	s.AddAttribute("refy", value)
	return s
}
