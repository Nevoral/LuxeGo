package svg

import (
	"LuxeGo/internal/lx"
)

//Marker - 
func Marker(tags ...lx.Content) *MarkerTag {
	return &MarkerTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "marker", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MarkerTag struct {
	*ComponentHtmlTag
}

//MarkerHeight - 
func (m *MarkerTag) MarkerHeight(value string) *MarkerTag {
	m.AddAttribute("markerheight", value)
	return m
}

//MarkerUnits - 
func (m *MarkerTag) MarkerUnits(value string) *MarkerTag {
	m.AddAttribute("markerunits", value)
	return m
}

//MarkerWidth - 
func (m *MarkerTag) MarkerWidth(value string) *MarkerTag {
	m.AddAttribute("markerwidth", value)
	return m
}

//Orient - 
func (m *MarkerTag) Orient(value string) *MarkerTag {
	m.AddAttribute("orient", value)
	return m
}

//PreserveAspectRatio - 
func (m *MarkerTag) PreserveAspectRatio(value string) *MarkerTag {
	m.AddAttribute("preserveaspectratio", value)
	return m
}

//RefX - 
func (m *MarkerTag) RefX(value string) *MarkerTag {
	m.AddAttribute("refx", value)
	return m
}

//RefY - 
func (m *MarkerTag) RefY(value string) *MarkerTag {
	m.AddAttribute("refy", value)
	return m
}

//ViewBox - 
func (m *MarkerTag) ViewBox(value string) *MarkerTag {
	m.AddAttribute("viewbox", value)
	return m
}
