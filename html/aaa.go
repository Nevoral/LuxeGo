package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
	"strings"
)

func (c *ComponentHtmlTag) AddAttribute(name, value string) {
	if c.Attributes == nil {
		panic(c)
	}
	*c.Attributes = append(*c.Attributes, LuxeGo.Attribute{
		Name:  name,
		Value: value,
	})
}

func (c *ComponentHtmlTag) Render(level int) string {
	var attributes string
	tab := strings.Repeat("\t", level)
	if c.Attributes != nil {
		attributes = c.Attributes.Render()
	}

	if c.Children == nil {
		if c.Attributes != nil {
			return fmt.Sprintf("%s<%s%s>", tab, c.Name, attributes)
		}
		if c.Name == "" {
			return fmt.Sprintf("%s%s", tab, c.Msg)
		}
		return fmt.Sprintf("<%s%s-->", c.Name, c.Msg)
	}

	var result strings.Builder
	for _, val := range *c.Children {
		result.WriteString(fmt.Sprintf("\n%s", val.Render(level+1)))
	}
	return fmt.Sprintf("%s<%s%s>%s\n%s</%s>", tab, c.Name, attributes, result.String(), tab, c.Name)
}
