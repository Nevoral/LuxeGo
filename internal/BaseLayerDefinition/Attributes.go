package base

import (
	"fmt"
	"slices"
	"strings"

	assert "github.com/Nevoral/LuxeGo/pkg/Asserts"
)

type Attributes interface {
	RegisterAttribute(name, value string, boolean bool)

	getAttributes() *attributes
	renderAttributeHtml() string
	attributeExist(name string) bool
}

type attributes struct {
	names   []string
	values  []string
	boolean []bool
}

func NewAttributes() Attributes {
	return new(attributes)
}

func (a *attributes) RegisterAttribute(name, value string, boolean bool) {
	assert.New(
		!a.attributeExist(name),
		fmt.Sprintf("Assert attribute already used in this tag: attribute name %s", name),
	)

	a.names = append(a.names, name)
	a.values = append(a.values, value)
	a.boolean = append(a.boolean, boolean)
}

func (a *attributes) getAttributes() *attributes {
	return a
}

func (a *attributes) renderAttributeHtml() string {
	assert.New(
		len(a.names) == len(a.boolean) || len(a.values) == len(a.boolean),
		"Assert: lenght of all 3 fields need to be same length.",
	)
	var attributes string
	for i := 0; i < len(a.names); i++ {
		if a.boolean[i] {
			attributes += fmt.Sprintf(" %s", a.names[i])
		} else {
			attributes += fmt.Sprintf(" %s=\"%s\"", a.names[i], strings.ReplaceAll(a.values[i], "%", "%%"))
		}
	}
	return attributes
}
func (a *attributes) attributeExist(name string) bool {
	return slices.Contains(a.names, name)
}

type AttributeType int

const (
	String AttributeType = iota
	Boolean
	Number
	Enum
	Time
	Percentage
	Color
)

// func (a AttributeType) String() string {
// 	return [...]string{"string", "boolean", "number", "enum", "time", "percentage", "color"}[a]
// }

// type attribute struct {
// 	name          string
// 	value         string
// 	attributeType AttributeType
// }

// func newAttribute(name, value string, attrType AttributeType) *attribute {
// 	return &attribute{name, value, attrType}
// }

// func (a *attribute) renderHtmlAttribute() string {
// 	if a.attributeType == Boolean {
// 		return fmt.Sprintf(" %s", a.name)
// 	} else {
// 		return fmt.Sprintf(" %s=\"%s\"", a.name, strings.ReplaceAll(a.value, "%", "%%"))
// 	}
// }

// func (a *attribute) renderLuxeGoAttribute() string {
// 	name := removeAndCamelCase(a.name, ":", "-")
// 	name = capitalizeFirst(name)
// 	if a.attributeType == Boolean {
// 		return fmt.Sprintf(".%s()", name)
// 	} else {
// 		return fmt.Sprintf(".%s(\"%s\")", name, strings.ReplaceAll(a.value, "%", "%%"))
// 	}
// }
