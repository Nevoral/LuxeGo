package lx

import (
	"fmt"
	"strings"
)

type WebComponent struct {
	Name       string
	Attributes *Attributes
	Msg        string
	Children   *[]Content
}

func (w *WebComponent) AddAttribute(name, value string) {
	if w.Attributes == nil {
		panic(w)
	}
	*w.Attributes = append(*w.Attributes, Attribute{
		Name:  name,
		Value: value,
	})
}

func (w *WebComponent) Render(level int) string {
	var attributes string
	if w.Attributes != nil {
		attributes = w.Attributes.Render()
	}

	if w.Children == nil {
		if w.Attributes != nil {
			return fmt.Sprintf("<%s%s>", w.Name, attributes)
		}
		if w.Name == "" {
			return fmt.Sprintf("%s", w.Msg)
		}
		return fmt.Sprintf("<%s%s-->", w.Name, w.Msg)
	}

	var result strings.Builder
	for _, val := range *w.Children {
		result.WriteString(fmt.Sprintf("\n%s", val.Render(level+1)))
	}
	return fmt.Sprintf("<%s%s>%s\n</%s>", w.Name, attributes, result.String(), w.Name)
}
