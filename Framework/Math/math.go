package math

import (
	"fmt"
	"reflect"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

type MathNode struct {
	base.Content
}

func Math(attributes *math, tags ...any) *MathNode {
	var children []base.Content
	for _, tag := range tags {
		switch t := tag.(type) {
		case string:
			children = append(children, stringValue(t))
		case *MathNode:
			children = append(children, t.Content)
		default:
			panic(fmt.Sprintf("assert \"%s\" tag doesn't belong into this namespace HTML", reflect.TypeOf(t).String()))
		}
	}
	return &MathNode{base.NewFullTagNode(
		"math",
		model.MATH,
		attributes.Attributes,
		&children,
	)}
}

func AttributesMath() *math {
	return &math{base.NewAttributes()}
}

type math struct {
	base.Attributes
}

func (m *math) Width(value string) *math {
	if value != "" {
		m.RegisterAttribute("width", value, false)
	}
	return m
}
