package html

import (
	"LuxeGo/internal/lx"
)

//Main - 
func Main(tags ...lx.Content) *MainTag {
	return &MainTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "main", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type MainTag struct {
	*ComponentTag
}
