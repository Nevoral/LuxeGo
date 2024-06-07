package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Progress -
func Progress(tags ...interface{}) *ProgressTag {
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
	return &ProgressTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "progress", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ProgressTag struct {
	*ComponentHtmlTag
}

// Max -
func (p *ProgressTag) Max(value string) *ProgressTag {
	p.AddAttribute("max", value)
	return p
}

// Value - Specifies the value of an <input> element.
func (p *ProgressTag) Value(value string) *ProgressTag {
	p.AddAttribute("value", value)
	return p
}
