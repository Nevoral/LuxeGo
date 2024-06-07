package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Var -
func Var(tags ...interface{}) *VarTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &VarTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "var", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type VarTag struct {
	*ComponentHtmlTag
}
