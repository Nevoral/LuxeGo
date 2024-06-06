package svg

import (
	"LuxeGo/internal/lx"
)

// Polyline -
func Polyline() *PolylineTag {
	return &PolylineTag{ComponentSvgTag: &ComponentSvgTag{Name: "polyline", Attributes: &lx.Attributes{}, Children: nil}}
}

type PolylineTag struct {
	*ComponentSvgTag
}

// Points -
func (p *PolylineTag) Points(value string) *PolylineTag {
	p.AddAttribute("points", value)
	return p
}

// PathLength -
func (p *PolylineTag) PathLength(value string) *PolylineTag {
	p.AddAttribute("pathlength", value)
	return p
}
