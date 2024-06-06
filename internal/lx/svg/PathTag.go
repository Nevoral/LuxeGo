package svg

import (
	"LuxeGo/internal/lx"
)

// Path -
func Path() *PathTag {
	return &PathTag{ComponentSvgTag: &ComponentSvgTag{Name: "path", Attributes: &lx.Attributes{}, Children: nil}}
}

type PathTag struct {
	*ComponentSvgTag
}

// D -
func (p *PathTag) D(value string) *PathTag {
	p.AddAttribute("d", value)
	return p
}

// PathLength -
func (p *PathTag) PathLength(value string) *PathTag {
	p.AddAttribute("pathlength", value)
	return p
}
