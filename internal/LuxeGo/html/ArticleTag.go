package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Article -
func Article(tags ...LuxeGo.Content) *ArticleTag {
	return &ArticleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "article", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ArticleTag struct {
	*ComponentHtmlTag
}
