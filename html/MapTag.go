package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Map -
func Map(tags ...interface{}) *MapTag {
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
	return &MapTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "map", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type MapTag struct {
	*ComponentHtmlTag
}

// Name -
func (m *MapTag) Name(value string) *MapTag {
	m.AddAttribute("name", value)
	return m
}
