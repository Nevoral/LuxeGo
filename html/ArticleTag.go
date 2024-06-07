package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Article -
func Article(tags ...interface{}) *ArticleTag {
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
	return &ArticleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "article", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type ArticleTag struct {
	*ComponentHtmlTag
}
