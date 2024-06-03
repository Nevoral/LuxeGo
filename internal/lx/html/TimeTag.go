package html

import (
	"LuxeGo/internal/lx"
)

//Time - 
func Time(tags ...lx.Content) *TimeTag {
	return &TimeTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "time", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TimeTag struct {
	*ComponentTag
}

//Datetime - Specifies the date and time.
func (t *TimeTag) Datetime(value string) *TimeTag {
	t.AddAttribute("datetime", value)
	return t
}
