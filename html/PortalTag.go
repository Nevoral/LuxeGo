package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Portal -
func Portal(tags ...interface{}) *PortalTag {
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
	return &PortalTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "portal", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type PortalTag struct {
	*ComponentHtmlTag
}

// Referrerpolicy -
func (p *PortalTag) Referrerpolicy(value string) *PortalTag {
	p.AddAttribute("referrerpolicy", value)
	return p
}

// Src - Specifies the URL of an image.
func (p *PortalTag) Src(value string) *PortalTag {
	p.AddAttribute("src", value)
	return p
}
