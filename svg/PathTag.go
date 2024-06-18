package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Path -
func Path() *PathTag {
	return &PathTag{ComponentSvgTag: &ComponentSvgTag{Name: "path", Attributes: &LuxeGo.Attributes{}, Children: nil}}
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

// FillRule -
func (p *PathTag) FillRule(value string) *PathTag {
	p.AddAttribute("fill-rule", value)
	return p
}

// ClipRule -
func (p *PathTag) ClipRule(value string) *PathTag {
	p.AddAttribute("clip-rule", value)
	return p
}

// Fill -
func (p *PathTag) Fill(value string) *PathTag {
	p.AddAttribute("fill", value)
	return p
}
