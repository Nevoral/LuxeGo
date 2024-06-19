package html

import (
	"context"
	"fmt"
	"github.com/Nevoral/LuxeGo"
	"io"
	"strings"
)

// FreeStr -
func FreeStr(msg string) *ComponentHtmlTag {
	return &ComponentHtmlTag{Name: "", Attributes: nil, Msg: msg, Children: nil}
}

// Comment -
func Comment(comment string) *ComponentHtmlTag {
	return &ComponentHtmlTag{Name: "!--", Attributes: nil, Msg: comment, Children: nil}
}

func (c *ComponentHtmlTag) CustomAtr(name, value string) *ComponentHtmlTag {
	c.AddAttribute(name, value)
	return c
}

func (c *ComponentHtmlTag) AddAttribute(name, value string) {
	if c.Attributes == nil {
		panic(c)
	}
	*c.Attributes = append(*c.Attributes, LuxeGo.Attribute{
		Name:  name,
		Value: value,
	})
}

func (c *ComponentHtmlTag) Render(ctx context.Context, writer io.Writer) error {
	var attributes string
	if c.Attributes != nil {
		attributes = c.Attributes.Render()
	}

	if c.Children == nil {
		if c.Attributes != nil {
			_, err := fmt.Fprintf(writer, "<%s%s>", c.Name, attributes)
			return err
		}
		if c.Name == "" {
			_, err := fmt.Fprintf(writer, "%s", c.Msg)
			return err
		}
		_, err := fmt.Fprintf(writer, "<%s%s-->", c.Name, c.Msg)
		return err
	}

	var result strings.Builder
	for _, val := range *c.Children {
		result.WriteString("\n")
		err := val.Render(ctx, &result)
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(writer, "<%s%s>%s\n</%s>", c.Name, attributes, result.String(), c.Name)
	return err
}

/*
func (c *ComponentHtmlTag) RenderString(level int) string {
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
		result.WriteString(fmt.Sprintf("\n%s", val.RenderString(level+1)))
	}
	return fmt.Sprintf("%s<%s%s>%s\n%s</%s>", tab, c.Name, attributes, result.String(), tab, c.Name)
}
*/
