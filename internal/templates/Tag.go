package templates

import (
	"fmt"
	"strings"
)

func TagCaller(name, comment, typeComponent string, isBool bool) string {
	if isBool {
		return fmt.Sprintf(`//%s - %s
func %s() *%sTag {
	return &%sTag{Component%sTag: &Component%sTag{Name: "%s", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

%s
`, name, comment, name, name, name, typeComponent, typeComponent, strings.ToLower(name), TagStruct(name, typeComponent))
	}
	return fmt.Sprintf(`//%s - %s
func %s(tags ...interface{}) *%sTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %%T", v))
		}
	}
	return &%sTag{Component%sTag: &Component%sTag{Name: "%s", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

%s
`, name, comment, name, name, name, typeComponent, typeComponent, strings.ToLower(name), TagStruct(name, typeComponent))
}

func TagStruct(name, typeComponent string) string {
	return fmt.Sprintf(`type %sTag struct {
	*Component%sTag
}`, name, typeComponent)
}

func TagMethod(tagName, atrName, comment string, isBool bool) string {
	name := RemoveAndCamelCase(atrName, "-")
	name = RemoveAndCamelCase(name, ":")
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
