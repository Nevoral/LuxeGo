package html

// Comment -
func Comment(comment string) *CommentTag {
	return &CommentTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!--", Attributes: nil, Msg: comment, Children: nil}}
}

type CommentTag struct {
	*ComponentHtmlTag
}
