package html

import (
	"LuxeGo/internal/lx"
)

//Summary - 
func Summary(tags ...lx.Content) *SummaryTag {
	return &SummaryTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "summary", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SummaryTag struct {
	*ComponentTag
}
