package html

import (
	"fmt"
	"strings"
)

func TagCaller(name, comment string, isBool bool) string {
	if isBool {
		return fmt.Sprintf(`//%s - %s
func %s() *%sTag {
	return &%sTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "%s", Attributes: &lx.Attributes{}, Children: nil}}}
}

%s
`, name, comment, name, name, name, strings.ToLower(name), TagStruct(name))
	}
	return fmt.Sprintf(`//%s - %s
func %s(tags ...lx.Content) *%sTag {
	return &%sTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "%s", Attributes: &lx.Attributes{}, Children: &tags}}}
}

%s
`, name, comment, name, name, name, strings.ToLower(name), TagStruct(name))
}

func TagStruct(name string) string {
	return fmt.Sprintf(`type %sTag struct {
	*ComponentTag
}`, name)
}

func TagMethod(tagName, atrName, comment string, isBool bool) string {
	firstLetter := string(strings.ToLower(tagName)[0])
	if isBool {
		return fmt.Sprintf(`
//%s - %s
func (%s *%sTag) %s() *%sTag {
	%s.AddAttribute("%s", "")
	return %s
}
`, atrName, comment, firstLetter, tagName, atrName, tagName, firstLetter, strings.ToLower(atrName), firstLetter)
	}
	return fmt.Sprintf(`
//%s - %s
func (%s *%sTag) %s(value string) *%sTag {
	%s.AddAttribute("%s", value)
	return %s
}
`, atrName, comment, firstLetter, tagName, atrName, tagName, firstLetter, strings.ToLower(atrName), firstLetter)
}
