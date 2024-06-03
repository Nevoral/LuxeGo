package html

import (
	"LuxeGo/internal/lx"
)

//Article - 
func Article(tags ...lx.Content) *ArticleTag {
	return &ArticleTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "article", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type ArticleTag struct {
	*ComponentTag
}
