package html

import (
	"LuxeGo/internal/lx"
)

//Article - 
func Article(tags ...lx.Content) *ArticleTag {
	return &ArticleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "article", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ArticleTag struct {
	*ComponentHtmlTag
}
