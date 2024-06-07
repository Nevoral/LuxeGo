package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Dialog -
func Dialog(tags ...LuxeGo.Content) *DialogTag {
	return &DialogTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dialog", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DialogTag struct {
	*ComponentHtmlTag
}

// Open -
func (d *DialogTag) Open() *DialogTag {
	d.AddAttribute("open", "")
	return d
}
