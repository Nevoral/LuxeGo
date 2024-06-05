package svg

import (
	"LuxeGo/internal/lx"
)

//FeFlood - 
func FeFlood() *FeFloodTag {
	return &FeFloodTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "feflood", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeFloodTag struct {
	*ComponentHtmlTag
}

//Flood-color - 
func (f *FeFloodTag) Flood-color(value string) *FeFloodTag {
	f.AddAttribute("flood-color", value)
	return f
}

//Flood-opacity - 
func (f *FeFloodTag) Flood-opacity(value string) *FeFloodTag {
	f.AddAttribute("flood-opacity", value)
	return f
}
