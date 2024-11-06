package html

import (
	"fmt"
	"reflect"
	"slices"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"

	assert "github.com/Nevoral/LuxeGo/pkg/Asserts"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

func A(attributes *a, tags ...any) base.Content {
	var (
		supportedNodes = []string{"a", "area", "svg", "math", "div", "p"}
		children       []base.Content
	)
	for _, tag := range tags {
		switch t := tag.(type) {
		case string:
			children = append(children, stringValue(t))
		case base.Content:
			assert.New(
				model.HTML == t.GetNamespace() || t.GetName() == "svg" || t.GetName() == "math",
				fmt.Sprintf("Assert %s tag isn't suported inside {{.TagName}} tag.", t.GetName()),
			)
			assert.New(slices.Contains(supportedNodes, t.GetName()), fmt.Sprintf("Assert %s tag isn't suported inside {{.TagName}} tag.", t.GetName()))
			children = append(children, t)
		default:
			panic(fmt.Sprintf("assert \"%s\" type isn't supported.", reflect.TypeOf(t).String()))
		}
	}
	return base.NewFullTagNode(
		"a",
		model.HTML,
		attributes.Attributes,
		&children,
	)
}

func AttributesOfA() *a {
	return &a{base.NewAttributes()}
}

type a struct {
	base.Attributes
}

func (a *a) CustomAttribute(name, value string, boolean bool) *a {
	a.RegisterAttribute(name, value, boolean)
	return a
}

func (a *a) Class(value string) *a {
	if value != "" {
		a.RegisterAttribute("class", value, false)
	}
	return a
}

func (a *a) Id(value string) *a {
	if value != "" {
		a.RegisterAttribute("id", value, false)
	}
	return a
}
