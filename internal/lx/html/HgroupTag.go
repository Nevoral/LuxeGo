package html

import (
	"LuxeGo/internal/lx"
)

//Hgroup - 
func Hgroup(tags ...lx.Content) *HgroupTag {
	return &HgroupTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "hgroup", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type HgroupTag struct {
	*ComponentTag
}
