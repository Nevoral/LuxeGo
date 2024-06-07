package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeSpotLight -
func FeSpotLight() *FeSpotLightTag {
	return &FeSpotLightTag{ComponentSvgTag: &ComponentSvgTag{Name: "fespotlight", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeSpotLightTag struct {
	*ComponentSvgTag
}

// X -
func (f *FeSpotLightTag) X(value string) *FeSpotLightTag {
	f.AddAttribute("x", value)
	return f
}

// Y -
func (f *FeSpotLightTag) Y(value string) *FeSpotLightTag {
	f.AddAttribute("y", value)
	return f
}

// Z -
func (f *FeSpotLightTag) Z(value string) *FeSpotLightTag {
	f.AddAttribute("z", value)
	return f
}

// PointsAtX -
func (f *FeSpotLightTag) PointsAtX(value string) *FeSpotLightTag {
	f.AddAttribute("pointsatx", value)
	return f
}

// PointsAtY -
func (f *FeSpotLightTag) PointsAtY(value string) *FeSpotLightTag {
	f.AddAttribute("pointsaty", value)
	return f
}

// PointsAtZ -
func (f *FeSpotLightTag) PointsAtZ(value string) *FeSpotLightTag {
	f.AddAttribute("pointsatz", value)
	return f
}

// SpecularExponent -
func (f *FeSpotLightTag) SpecularExponent(value string) *FeSpotLightTag {
	f.AddAttribute("specularexponent", value)
	return f
}

// LimitingConeAngle -
func (f *FeSpotLightTag) LimitingConeAngle(value string) *FeSpotLightTag {
	f.AddAttribute("limitingconeangle", value)
	return f
}
