package svg

import (
	"LuxeGo/internal/lx"
)

//RadialGradient - 
func RadialGradient(tags ...lx.Content) *RadialGradientTag {
	return &RadialGradientTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "radialgradient", Attributes: &lx.Attributes{}, Children: &tags}}
}

type RadialGradientTag struct {
	*ComponentHtmlTag
}

//Cx - 
func (r *RadialGradientTag) Cx(value string) *RadialGradientTag {
	r.AddAttribute("cx", value)
	return r
}

//Cy - 
func (r *RadialGradientTag) Cy(value string) *RadialGradientTag {
	r.AddAttribute("cy", value)
	return r
}

//R - 
func (r *RadialGradientTag) R(value string) *RadialGradientTag {
	r.AddAttribute("r", value)
	return r
}

//Fx - 
func (r *RadialGradientTag) Fx(value string) *RadialGradientTag {
	r.AddAttribute("fx", value)
	return r
}

//Fy - 
func (r *RadialGradientTag) Fy(value string) *RadialGradientTag {
	r.AddAttribute("fy", value)
	return r
}

//Fr - 
func (r *RadialGradientTag) Fr(value string) *RadialGradientTag {
	r.AddAttribute("fr", value)
	return r
}

//GradientUnits - 
func (r *RadialGradientTag) GradientUnits(value string) *RadialGradientTag {
	r.AddAttribute("gradientunits", value)
	return r
}

//GradientTransform - 
func (r *RadialGradientTag) GradientTransform(value string) *RadialGradientTag {
	r.AddAttribute("gradienttransform", value)
	return r
}

//SpreadMethod - 
func (r *RadialGradientTag) SpreadMethod(value string) *RadialGradientTag {
	r.AddAttribute("spreadmethod", value)
	return r
}
