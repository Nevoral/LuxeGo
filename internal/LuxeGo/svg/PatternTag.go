package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// Pattern -
func Pattern(tags ...LuxeGo.Content) *PatternTag {
	return &PatternTag{ComponentSvgTag: &ComponentSvgTag{Name: "pattern", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type PatternTag struct {
	*ComponentSvgTag
}

// Height -
func (p *PatternTag) Height(value string) *PatternTag {
	p.AddAttribute("height", value)
	return p
}

// PatternContentUnits -
func (p *PatternTag) PatternContentUnits(value string) *PatternTag {
	p.AddAttribute("patterncontentunits", value)
	return p
}

// PatternTransform -
func (p *PatternTag) PatternTransform(value string) *PatternTag {
	p.AddAttribute("patterntransform", value)
	return p
}

// PatternUnits -
func (p *PatternTag) PatternUnits(value string) *PatternTag {
	p.AddAttribute("patternunits", value)
	return p
}

// PreserveAspectRatio -
func (p *PatternTag) PreserveAspectRatio(value string) *PatternTag {
	p.AddAttribute("preserveaspectratio", value)
	return p
}

// ViewBox -
func (p *PatternTag) ViewBox(value string) *PatternTag {
	p.AddAttribute("viewbox", value)
	return p
}

// Width -
func (p *PatternTag) Width(value string) *PatternTag {
	p.AddAttribute("width", value)
	return p
}

// X -
func (p *PatternTag) X(value string) *PatternTag {
	p.AddAttribute("x", value)
	return p
}

// Y -
func (p *PatternTag) Y(value string) *PatternTag {
	p.AddAttribute("y", value)
	return p
}
