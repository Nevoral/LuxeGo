package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Mpath -
func Mpath(tags ...interface{}) *MpathTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &MpathTag{ComponentSvgTag: &ComponentSvgTag{Name: "mpath", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MpathTag struct {
	*ComponentSvgTag
}

// Href -
func (m *MpathTag) Href(value string) *MpathTag {
	m.AddAttribute("href", value)
	return m
}
