package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Symbol -
func Symbol(tags ...LuxeGo.Content) *SymbolTag {
	return &SymbolTag{ComponentSvgTag: &ComponentSvgTag{Name: "symbol", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
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
