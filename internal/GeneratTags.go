package internal

import (
	"LuxeGo/configuration"
	html "LuxeGo/internal/templates/html"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func GenerateHtmlTags() {
	page := `package html

import (
	"LuxeGo/internal/lx"
)

%s`
	for key, val := range configuration.HtmlAtrTable {
		if key == "!DOCTYPE" || key == "!--" {
			continue
		}
		content := ""
		tagName := capitalizeFirst(key)
		content += html.TagCaller(tagName, "", slices.Contains(configuration.SelfClosingHtmlTag, key))
		for _, atr := range val {
			content += html.TagMethod(tagName, capitalizeFirst(atr), configuration.SpecificAtrTable[atr], slices.Contains(configuration.BoolAtr, atr))
		}
		CreateFile(fmt.Sprintf("internal/lx/html/%sTag.go", tagName), fmt.Sprintf(page, content))
	}
}
func GenerataGlobalAtr() {
	page := `package html

import "LuxeGo/internal/lx"

type ComponentTag struct {
	*lx.WebComponent
}

%s`
	content := ""
	for atr, comment := range configuration.GlobalAtrTable {
		content += html.TagMethod("Component", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
	}
	for atr, comment := range configuration.GlobalEventAtrTable {
		content += html.TagMethod("Component", capitalizeFirst(atr), comment, slices.Contains(configuration.BoolAtr, atr))
	}
	CreateFile(fmt.Sprintf("internal/lx/html/a_Component.go"), fmt.Sprintf(page, content))
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
