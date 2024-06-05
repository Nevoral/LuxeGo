package html

import (
	"LuxeGo/internal/lx"
)

//Summary - 
func Summary(tags ...lx.Content) *SummaryTag {
	return &SummaryTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "summary", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SummaryTag struct {
	*ComponentHtmlTag
}
