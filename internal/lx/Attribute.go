package lx

import "fmt"

type Attribute struct {
	Name  string
	Value string
}

type Attributes []Attribute

func (a *Attributes) Render() string {
	var result string
	for _, val := range *a {
		if val.Value != "" {
			result += fmt.Sprintf(" %s=\"%s\"", val.Name, val.Value)
		} else {
			result += fmt.Sprintf(" %s", val.Name)
		}
	}
	return result
}
