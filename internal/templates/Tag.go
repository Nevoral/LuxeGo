package templates

import (
	"fmt"
	"strings"
)

func TagCaller(name, comment string, isBool bool) string {
	if isBool {
		return fmt.Sprintf(`//%s - %s
func %s() *%sTag {
	return &%sTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "%s", Attributes: &lx.Attributes{}, Children: nil}}
}

%s
`, name, comment, name, name, name, strings.ToLower(name), TagStruct(name))
	}
	return fmt.Sprintf(`//%s - %s
func %s(tags ...lx.Content) *%sTag {
	return &%sTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "%s", Attributes: &lx.Attributes{}, Children: &tags}}
}

%s
`, name, comment, name, name, name, strings.ToLower(name), TagStruct(name))
}

func TagStruct(name string) string {
	return fmt.Sprintf(`type %sTag struct {
	*ComponentHtmlTag
}`, name)
}

func TagMethod(tagName, atrName, comment string, isBool bool) string {
	name := RemoveAndCamelCase(atrName, "-")
	name = RemoveAndCamelCase(atrName, ":")
	firstLetter := string(strings.ToLower(tagName)[0])
	if isBool {
		return fmt.Sprintf(`
//%s - %s
func (%s *%sTag) %s() *%sTag {
	%s.AddAttribute("%s", "")
	return %s
}
`, name, comment, firstLetter, tagName, name, tagName, firstLetter, strings.ToLower(atrName), firstLetter)
	}
	return fmt.Sprintf(`
//%s - %s
func (%s *%sTag) %s(value string) *%sTag {
	%s.AddAttribute("%s", value)
	return %s
}
`, name, comment, firstLetter, tagName, name, tagName, firstLetter, strings.ToLower(atrName), firstLetter)
}

func RemoveAndCamelCase(atrName, char string) string {
	name := atrName
	if strings.Contains(atrName, char) {
		index := strings.Index(atrName, char)
		name = strings.Replace(name, char, "", -1)
		name = name[:index] + strings.ToUpper(string(name[index])) + name[index+1:]
	}
	return name
}
