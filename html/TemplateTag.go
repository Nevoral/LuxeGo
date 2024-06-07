package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Template -
func Template(tags ...interface{}) *TemplateTag {
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
	return &TemplateTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "template", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TemplateTag struct {
	*ComponentHtmlTag
}
