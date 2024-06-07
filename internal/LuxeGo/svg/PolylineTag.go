package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// Polyline -
func Polyline() *PolylineTag {
	return &PolylineTag{ComponentSvgTag: &ComponentSvgTag{Name: "polyline", Attributes: &LuxeGo.Attributes{}, Children: nil}}
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
