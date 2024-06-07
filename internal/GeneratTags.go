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
	%s"github.com/Nevoral/LuxeGo"
)

`
	val := configuration.HtmlAtrTable[key]
	content := ""
	tagName := capitalizeFirst(key)
	isSelfClosing := slices.Contains(configuration.SelfClosingHtmlTag, key)
	switch key {
	case "!DOCTYPE":
		imports = fmt.Sprintf(imports, "")
		tagName = "Doctype"
		content += fmt.Sprintf(`//DOCTYPE - %s
func DOCTYPE() *DoctypeTag {
	return &DoctypeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "!DOCTYPE", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

%s
`, "", tmpl.TagStruct("Doctype", "Html"))
	default:
		if isSelfClosing {
			imports = fmt.Sprintf(imports, "")
		} else {
			imports = fmt.Sprintf(imports, "\"fmt\"\n\t")
		}
		content += tmpl.TagCaller(tagName, "", "Html", isSelfClosing)
	}
	for _, atr := range val {
		content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SpecificAtrTable[atr], slices.Contains(configuration.BoolAtr, atr))
	}
	CreateFile(fmt.Sprintf("html/%sTag.go", tagName), fmt.Sprintf(page, imports, content))
	return
}

func GenerateHtmlTags() {
	for key := range configuration.HtmlAtrTable {
		GenerateHtmlTag(key)
	}
}

func GenerataGlobalAtr() {
	page := `package html

import (
	"fmt"
	"slices"
	"github.com/Nevoral/LuxeGo"
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
	var a string
	if name == "" {
		a = "aria"
	} else {	
		a = fmt.Sprintf("aria-%%s", name)
	}
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
	var a string
	if name == "" {
		a = "data"
	} else {	
		a = fmt.Sprintf("data-%%s", name)
	}
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
	%s"github.com/Nevoral/LuxeGo"
)

%s`
	val := configuration.SvgAtrTable[key]
	content := ""
	tagName := capitalizeFirst(key)
	isSelfClosing := slices.Contains(configuration.SvgSelfClosing, key)
	content += tmpl.TagCaller(tagName, "", "Svg", isSelfClosing)
	for _, atr := range val {
		content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SvgSoecificAtr[atr], slices.Contains(configuration.BoolSvgAtr, atr))
	}
	if isSelfClosing {
		CreateFile(fmt.Sprintf("svg/%sTag.go", tagName), fmt.Sprintf(page, "", content))
		return
	}
	CreateFile(fmt.Sprintf("svg/%sTag.go", tagName), fmt.Sprintf(page, "\"fmt\"\n\t", content))
	return
}

func GenerateSvgTags() {
	for key := range configuration.SvgAtrTable {
		GenerateSvgTag(key)
	}
}

func GenerataGlobalAtrSvg() {
	page := `package svg

import (
	"fmt"
	"slices"
	"github.com/Nevoral/LuxeGo"
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
	var a string
	if name == "" {
		a = "aria"
	} else {	
		a = fmt.Sprintf("aria-%%s", name)
	}
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
	var a string
	if name == "" {
		a = "data"
	} else {	
		a = fmt.Sprintf("data-%%s", name)
	}
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
