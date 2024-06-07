package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Summary -
func Summary(tags ...LuxeGo.Content) *SummaryTag {
	return &SummaryTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "summary", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SummaryTag struct {
	*ComponentHtmlTag
}
