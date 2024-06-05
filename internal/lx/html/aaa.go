package html

import (
	"LuxeGo/internal/lx"
	"fmt"
	"strings"
)

func (c *ComponentHtmlTag) AddAttribute(name, value string) {
	if c.Attributes == nil {
		panic(c)
	}
	*c.Attributes = append(*c.Attributes, lx.Attribute{
		Name:  name,
		Value: value,
	})
}

func (c *ComponentHtmlTag) Render() string {
	var attributes string
	if c.Attributes != nil {
		attributes = c.Attributes.Render()
	}

	if c.Children == nil {
		if c.Attributes != nil {
			return fmt.Sprintf("<%s%s>", c.Name, attributes)
		}
		if c.Name == "" {
			return fmt.Sprintf("%s", c.Msg)
		}
		return fmt.Sprintf("<%s%s-->", c.Name, c.Msg)
	}

	var result strings.Builder
	for _, val := range *c.Children {
		result.WriteString(fmt.Sprintf("\n%s", val.Render()))
	}
	return fmt.Sprintf("<%s%s>%s\n</%s>", c.Name, attributes, result.String(), c.Name)
}
