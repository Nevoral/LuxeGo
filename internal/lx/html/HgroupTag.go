package html

import (
	"LuxeGo/internal/lx"
)

// Hgroup -
func Hgroup(tags ...lx.Content) *HgroupTag {
	return &HgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "hgroup", Attributes: &lx.Attributes{}, Children: &tags}}
}

type HgroupTag struct {
	*ComponentHtmlTag
}
