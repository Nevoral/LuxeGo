package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Dialog -
func Dialog(tags ...interface{}) *DialogTag {
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
	return &DialogTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dialog", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DialogTag struct {
	*ComponentHtmlTag
}

// Open -
func (d *DialogTag) Open() *DialogTag {
	d.AddAttribute("open", "")
	return d
}
