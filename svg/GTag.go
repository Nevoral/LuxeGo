package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// G -
func G(tags ...interface{}) *GTag {
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
	return &GTag{ComponentSvgTag: &ComponentSvgTag{Name: "g", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type GTag struct {
	*ComponentSvgTag
}

// StrokeWidth -
func (g *GTag) StrokeWidth(value string) *GTag {
	g.AddAttribute("stroke-width", value)
	return g
}

// StrokeLinecap -
func (g *GTag) StrokeLinecap(value string) *GTag {
	g.AddAttribute("stroke-linecap", value)
	return g
}

// StrokeLinejoin -
func (g *GTag) StrokeLinejoin(value string) *GTag {
	g.AddAttribute("stroke-linejoin", value)
	return g
}
