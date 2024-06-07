package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Time -
func Time(tags ...LuxeGo.Content) *TimeTag {
	return &TimeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "time", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TimeTag struct {
	*ComponentHtmlTag
}

// Datetime - Specifies the date and time.
func (t *TimeTag) Datetime(value string) *TimeTag {
	t.AddAttribute("datetime", value)
	return t
}
