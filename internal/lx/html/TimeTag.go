package html

import (
	"LuxeGo/internal/lx"
)

// Time -
func Time(tags ...lx.Content) *TimeTag {
	return &TimeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "time", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TimeTag struct {
	*ComponentHtmlTag
}

// Datetime - Specifies the date and time.
func (t *TimeTag) Datetime(value string) *TimeTag {
	t.AddAttribute("datetime", value)
	return t
}
