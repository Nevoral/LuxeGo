package html

// FreeStr -
func FreeStr(msg string) *FreeStrTag {
	return &FreeStrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "", Attributes: nil, Msg: msg, Children: nil}}
}

type FreeStrTag struct {
	*ComponentHtmlTag
}
