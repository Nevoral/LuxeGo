package svg

import (
	"LuxeGo/internal/lx"
)

//FeDistantLight - 
func FeDistantLight() *FeDistantLightTag {
	return &FeDistantLightTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fedistantlight", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeDistantLightTag struct {
	*ComponentHtmlTag
}

//Azimuth - 
func (f *FeDistantLightTag) Azimuth(value string) *FeDistantLightTag {
	f.AddAttribute("azimuth", value)
	return f
}

//Elevation - 
func (f *FeDistantLightTag) Elevation(value string) *FeDistantLightTag {
	f.AddAttribute("elevation", value)
	return f
}
