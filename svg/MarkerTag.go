package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Marker -
func Marker(tags ...interface{}) *MarkerTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &MarkerTag{ComponentSvgTag: &ComponentSvgTag{Name: "marker", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MarkerTag struct {
	*ComponentSvgTag
}

// MarkerHeight -
func (m *MarkerTag) MarkerHeight(value string) *MarkerTag {
	m.AddAttribute("markerheight", value)
	return m
}

// MarkerUnits -
func (m *MarkerTag) MarkerUnits(value string) *MarkerTag {
	m.AddAttribute("markerunits", value)
	return m
}

// MarkerWidth -
func (m *MarkerTag) MarkerWidth(value string) *MarkerTag {
	m.AddAttribute("markerwidth", value)
	return m
}

// Orient -
func (m *MarkerTag) Orient(value string) *MarkerTag {
	m.AddAttribute("orient", value)
	return m
}

// PreserveAspectRatio -
func (m *MarkerTag) PreserveAspectRatio(value string) *MarkerTag {
	m.AddAttribute("preserveaspectratio", value)
	return m
}

// RefX -
func (m *MarkerTag) RefX(value string) *MarkerTag {
	m.AddAttribute("refx", value)
	return m
}

// RefY -
func (m *MarkerTag) RefY(value string) *MarkerTag {
	m.AddAttribute("refy", value)
	return m
}

// ViewBox -
func (m *MarkerTag) ViewBox(value string) *MarkerTag {
	m.AddAttribute("viewbox", value)
	return m
}
