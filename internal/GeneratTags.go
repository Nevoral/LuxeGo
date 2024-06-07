package internal

import (
	"fmt"
	"github.com/Nevoral/LuxeGo/configuration"
	tmpl "github.com/Nevoral/LuxeGo/internal/templates"
	"os"
	"slices"
	"unicode"
)

func GenerateHtmlTag(key string) {
	page := `package html
%s
%s`
	imports := `
import (
	"github.com/Nevoral/LuxeGo"
)

`
	val := configuration.HtmlAtrTable[key]
	content := ""
	tagName := capitalizeFirst(key)
	switch key {
	case "!DOCTYPE":
		tagName = "Doctype"
		content += fmt.Sprintf(`//DOCTYPE - %s
func DOCTYPE() *DoctypeTag {
	return &DoctypeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!DOCTYPE", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

%s
`, "", tmpl.TagStruct("Doctype", "Html"))
	case "!--":
		tagName = "Comment"
		imports = ""
		content += fmt.Sprintf(`//Comment - %s
func Comment(comment string) *CommentTag {
	return &CommentTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!--", Attributes: nil, Msg: comment, Children: nil}}
}

%s
`, "", tmpl.TagStruct("Comment", "Html"))
	case "":
		tagName = "FreeStr"
		imports = ""
		content += fmt.Sprintf(`//FreeStr - %s
func FreeStr(msg string) *FreeStrTag {
	return &FreeStrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "", Attributes: nil,  Msg: msg, Children: nil}}
}

%s
`, "", tmpl.TagStruct("FreeStr", "Html"))
	default:
		content += tmpl.TagCaller(tagName, "", "Html", slices.Contains(configuration.SelfClosingHtmlTag, key))
	}
	for _, atr := range val {
		content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SpecificAtrTable[atr], slices.Contains(configuration.BoolAtr, atr))
	}
	CreateFile(fmt.Sprintf("html/%sTag.go", tagName), fmt.Sprintf(page, imports, content))
}

func GenerateHtmlTags() {
	for key := range configuration.HtmlAtrTable {
		GenerateHtmlTag(key)
	}
}

func GenerataGlobalAtr() {
	page := `package html

import (
	"github.com/Nevoral/LuxeGo"
	"fmt"
	"slices"
)

type ComponentHtmlTag struct {
	Name       string
	Attributes *LuxeGo.Attributes
	Msg        string
	Children   *[]LuxeGo.Content
}
%s`
	content := ""
	for atr, comment := range configuration.GlobalAtrTable {
		switch atr {
		case "aria":
			content += fmt.Sprintf(`
//Aria - %s
func (c *ComponentHtmlTag) Aria(name, value string) *ComponentHtmlTag {
	a := fmt.Sprintf("aria-%%s", name)
	if slices.Contains(LuxeGo.GlobAriaName, name) {
		c.AddAttribute(a, value)
	} else {
	//TODO: Specific aria attributes need to be check before adding to the map of attributes
		c.AddAttribute(a, value)
	}
	return c
}
`, comment)
		case "data":
			content += fmt.Sprintf(`
//Data - %s
func (c *ComponentHtmlTag) Data(name, value string) *ComponentHtmlTag {
	a := fmt.Sprintf("data-%%s", name)
	c.AddAttribute(a, value)
	return c
}
`, comment)
		default:
			content += tmpl.TagMethod("ComponentHtml", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
		}
	}
	for atr, comment := range configuration.GlobalEventAtrTable {
		content += tmpl.TagMethod("ComponentHtml", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
	}
	CreateFile(fmt.Sprintf("html/a_Component.go"), fmt.Sprintf(page, content))
}

func GenerateSvgTag(key string) {
	page := `package svg

import (
	"github.com/Nevoral/LuxeGo"
)

%s`
	val := configuration.SvgAtrTable[key]
	content := ""
	tagName := capitalizeFirst(key)
	content += tmpl.TagCaller(tagName, "", "Svg", slices.Contains(configuration.SvgSelfClosing, key))
	for _, atr := range val {
		content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SvgSoecificAtr[atr], slices.Contains(configuration.BoolSvgAtr, atr))
	}
	CreateFile(fmt.Sprintf("svg/%sTag.go", tagName), fmt.Sprintf(page, content))

}

func GenerateSvgTags() {
	for key := range configuration.SvgAtrTable {
		GenerateSvgTag(key)
	}
}

func GenerataGlobalAtrSvg() {
	page := `package svg

import (
	"github.com/Nevoral/LuxeGo"
	"fmt"
	"slices"
)

type ComponentSvgTag struct {
	Name       string
	Attributes *LuxeGo.Attributes
	Msg        string
	Children   *[]LuxeGo.Content
}
%s`
	content := ""
	for atr, comment := range configuration.SvgGlobalAtr {
		switch atr {
		case "aria":
			content += fmt.Sprintf(`
//Aria - %s
func (c *ComponentSvgTag) Aria(name, value string) *ComponentSvgTag {
	a := fmt.Sprintf("aria-%%s", name)
	if slices.Contains(LuxeGo.GlobAriaName, name) {
		c.AddAttribute(a, value)
	} else {
	//TODO: Specific aria attributes need to be check before adding to the map of attributes
		c.AddAttribute(a, value)
	}
	return c
}
`, comment)
		case "data":
			content += fmt.Sprintf(`
//Data - %s
func (c *ComponentSvgTag) Data(name, value string) *ComponentSvgTag {
	a := fmt.Sprintf("data-%%s", name)
	c.AddAttribute(a, value)
	return c
}
`, comment)
		default:
			content += tmpl.TagMethod("ComponentSvg", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolSvgAtr, atr))
		}
	}
	for atr, comment := range configuration.GlobalEventAtrTable {
		content += tmpl.TagMethod("ComponentSvg", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolSvgAtr, atr))
	}
	CreateFile(fmt.Sprintf("svg/a_Component.go"), fmt.Sprintf(page, content))
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func CreateFile(fileName, message string) {
	// Create a new file (truncates the file if it already exists)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.WriteString(message)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File '%s' created successfully.\n", fileName)
}
