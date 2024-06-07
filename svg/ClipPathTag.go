package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// ClipPath -
func ClipPath(tags ...interface{}) *ClipPathTag {
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
	return &ClipPathTag{ComponentSvgTag: &ComponentSvgTag{Name: "clippath", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ClipPathTag struct {
	*ComponentSvgTag
}

// ClipPathUnits -
func (c *ClipPathTag) ClipPathUnits(value string) *ClipPathTag {
	c.AddAttribute("clippathunits", value)
	return c
}

// Transform -
func (c *ClipPathTag) Transform(value string) *ClipPathTag {
	c.AddAttribute("transform", value)
	return c
}
