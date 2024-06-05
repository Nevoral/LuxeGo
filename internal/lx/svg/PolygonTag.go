package svg

import (
	"LuxeGo/internal/lx"
)

//Polygon - 
func Polygon() *PolygonTag {
	return &PolygonTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "polygon", Attributes: &lx.Attributes{}, Children: nil}}
}

type PolygonTag struct {
	*ComponentHtmlTag
}

//Points - 
func (p *PolygonTag) Points(value string) *PolygonTag {
	p.AddAttribute("points", value)
	return p
}

//PathLength - 
func (p *PolygonTag) PathLength(value string) *PolygonTag {
	p.AddAttribute("pathlength", value)
	return p
}
