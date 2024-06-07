package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Hgroup -
func Hgroup(tags ...LuxeGo.Content) *HgroupTag {
	return &HgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hgroup", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type HgroupTag struct {
	*ComponentHtmlTag
}
