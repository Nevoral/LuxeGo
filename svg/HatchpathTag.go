package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Hatchpath -
func Hatchpath(tags ...interface{}) *HatchpathTag {
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
	return &HatchpathTag{ComponentSvgTag: &ComponentSvgTag{Name: "hatchpath", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type HatchpathTag struct {
	*ComponentSvgTag
}

// D -
func (h *HatchpathTag) D(value string) *HatchpathTag {
	h.AddAttribute("d", value)
	return h
}

// Offset -
func (h *HatchpathTag) Offset(value string) *HatchpathTag {
	h.AddAttribute("offset", value)
	return h
}

// StrokeWidth -
func (h *HatchpathTag) StrokeWidth(value string) *HatchpathTag {
	h.AddAttribute("strokewidth", value)
	return h
}

// PathLength -
func (h *HatchpathTag) PathLength(value string) *HatchpathTag {
	h.AddAttribute("pathlength", value)
	return h
}
