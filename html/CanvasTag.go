package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Canvas -
func Canvas(tags ...interface{}) *CanvasTag {
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
	return &CanvasTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "canvas", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type CanvasTag struct {
	*ComponentHtmlTag
}

// Height -
func (c *CanvasTag) Height(value string) *CanvasTag {
	c.AddAttribute("height", value)
	return c
}

// Width -
func (c *CanvasTag) Width(value string) *CanvasTag {
	c.AddAttribute("width", value)
	return c
}
