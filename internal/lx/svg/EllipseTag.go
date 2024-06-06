package svg

import (
	"LuxeGo/internal/lx"
)

// Ellipse -
func Ellipse() *EllipseTag {
	return &EllipseTag{ComponentSvgTag: &ComponentSvgTag{Name: "ellipse", Attributes: &lx.Attributes{}, Children: nil}}
}

type EllipseTag struct {
	*ComponentSvgTag
}

// Cx -
func (e *EllipseTag) Cx(value string) *EllipseTag {
	e.AddAttribute("cx", value)
	return e
}

// Cy -
func (e *EllipseTag) Cy(value string) *EllipseTag {
	e.AddAttribute("cy", value)
	return e
}

// Rx -
func (e *EllipseTag) Rx(value string) *EllipseTag {
	e.AddAttribute("rx", value)
	return e
}

// Ry -
func (e *EllipseTag) Ry(value string) *EllipseTag {
	e.AddAttribute("ry", value)
	return e
}

// PathLength -
func (e *EllipseTag) PathLength(value string) *EllipseTag {
	e.AddAttribute("pathlength", value)
	return e
}
