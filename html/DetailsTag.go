package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Details -
func Details(tags ...interface{}) *DetailsTag {
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
	return &DetailsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "details", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DetailsTag struct {
	*ComponentHtmlTag
}

// Open -
func (d *DetailsTag) Open() *DetailsTag {
	d.AddAttribute("open", "")
	return d
}
