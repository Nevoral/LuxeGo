package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Var -
func Var(tags ...LuxeGo.Content) *VarTag {
	return &VarTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "var", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type VarTag struct {
	*ComponentHtmlTag
}
