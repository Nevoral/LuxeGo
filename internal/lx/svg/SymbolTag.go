package svg

import (
	"LuxeGo/internal/lx"
)

//Symbol - 
func Symbol(tags ...lx.Content) *SymbolTag {
	return &SymbolTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "symbol", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SymbolTag struct {
	*ComponentHtmlTag
}

//ViewBox - 
func (s *SymbolTag) ViewBox(value string) *SymbolTag {
	s.AddAttribute("viewbox", value)
	return s
}

//PreserveAspectRatio - 
func (s *SymbolTag) PreserveAspectRatio(value string) *SymbolTag {
	s.AddAttribute("preserveaspectratio", value)
	return s
}

//RefX - 
func (s *SymbolTag) RefX(value string) *SymbolTag {
	s.AddAttribute("refx", value)
	return s
}

//RefY - 
func (s *SymbolTag) RefY(value string) *SymbolTag {
	s.AddAttribute("refy", value)
	return s
}
