package internal

import (
	"LuxeGo/configuration"
	tmpl "LuxeGo/internal/templates"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func GenerateHtmlTags() {
	page := `package html
%s
%s`
	for key, val := range configuration.HtmlAtrTable {
		imports := `
import (
	"LuxeGo/internal/lx"
)

`
		content := ""
		tagName := capitalizeFirst(key)
		switch key {
		case "!DOCTYPE":
			tagName = tagName[1:]
			content += fmt.Sprintf(`//%s - %s
func %s() *%sTag {
	return &%sTag{Component%sTag: &Component%sTag{Name: "%s", Attributes: &lx.Attributes{}, Children: nil}}
}

%s
`, tagName, "", tagName, tagName, tagName, "Html", "Html", key, tmpl.TagStruct(tagName, "Html"))
			break
		case "!--":
			tagName = "Comment"
			imports = ""
			content += fmt.Sprintf(`//%s - %s
func %s(comment string) *%sTag {
	return &%sTag{Component%sTag: &Component%sTag{Name: "%s", Attributes: nil, Msg: comment, Children: nil}}
}

%s
`, tagName, "", tagName, tagName, tagName, "Html", "Html", key, tmpl.TagStruct(tagName, "Html"))
			break
		case "":
			tagName = "FreeStr"
			imports = ""
			content += fmt.Sprintf(`//%s - %s
func %s(msg string) *%sTag {
	return &%sTag{Component%sTag: &Component%sTag{Name: "%s", Attributes: nil,  Msg: msg, Children: nil}}
}

%s
`, tagName, "", tagName, tagName, tagName, "Html", "Html", "", tmpl.TagStruct(tagName, "Html"))
			break
		default:
			content += tmpl.TagCaller(tagName, "", "Html", slices.Contains(configuration.SelfClosingHtmlTag, key))
		}
		for _, atr := range val {
			content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SpecificAtrTable[atr], slices.Contains(configuration.BoolAtr, atr))
		}
		CreateFile(fmt.Sprintf("internal/lx/html/%sTag.go", tagName), fmt.Sprintf(page, imports, content))
	}
}

func GenerataGlobalAtr() {
	page := `package html

import "LuxeGo/internal/lx"

type ComponentHtmlTag struct {
	Name       string
	Attributes *lx.Attributes
	Msg        string
	Children   *[]lx.Content
}
%s`
	content := ""
	for atr, comment := range configuration.GlobalAtrTable {
		content += tmpl.TagMethod("ComponentHtml", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
	}
	for atr, comment := range configuration.GlobalEventAtrTable {
		content += tmpl.TagMethod("ComponentHtml", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
	}
	CreateFile(fmt.Sprintf("internal/lx/html/a_Component.go"), fmt.Sprintf(page, content))
}

func GenerateSvgTags() {
	page := `package svg

import (
	"LuxeGo/internal/lx"
)

%s`
	for key, val := range configuration.SvgAtrTable {
		if key == "!DOCTYPE" || key == "!--" || key == "" {
			continue
		}
		content := ""
		tagName := capitalizeFirst(key)
		content += tmpl.TagCaller(tagName, "", "Svg", slices.Contains(configuration.SvgSelfClosing, key))
		for _, atr := range val {
			content += tmpl.TagMethod(tagName, capitalizeFirst(atr), configuration.SvgSoecificAtr[atr], slices.Contains(configuration.BoolSvgAtr, atr))
		}
		CreateFile(fmt.Sprintf("internal/lx/svg/%sTag.go", tagName), fmt.Sprintf(page, content))
	}
}

func GenerataGlobalAtrSvg() {
	page := `package svg

import "LuxeGo/internal/lx"

type ComponentSvgTag struct {
	Name       string
	Attributes *lx.Attributes
	Msg        string
	Children   *[]lx.Content
}
%s`
	content := ""
	for atr, comment := range configuration.SvgGlobalAtr {
		content += tmpl.TagMethod("ComponentSvg", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolSvgAtr, atr))
	}
	for atr, comment := range configuration.GlobalEventAtrTable {
		content += tmpl.TagMethod("ComponentSvg", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolSvgAtr, atr))
	}
	CreateFile(fmt.Sprintf("internal/lx/svg/a_Component.go"), fmt.Sprintf(page, content))
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
