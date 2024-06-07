package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// Polygon -
func Polygon() *PolygonTag {
	return &PolygonTag{ComponentSvgTag: &ComponentSvgTag{Name: "polygon", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type PolygonTag struct {
	*ComponentSvgTag
}

// Points -
func (p *PolygonTag) Points(value string) *PolygonTag {
	p.AddAttribute("points", value)
	return p
}

// PathLength -
func (p *PolygonTag) PathLength(value string) *PolygonTag {
	p.AddAttribute("pathlength", value)
	return p
}
