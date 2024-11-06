package svg

import (
	"fmt"
	"reflect"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

type SvgNode struct {
	base.Content
}

func Svg(attributes *svg, tags ...any) *SvgNode {
	var children []base.Content
	for _, tag := range tags {
		switch t := tag.(type) {
		case string:
			children = append(children, stringValue(t))
		case *SvgNode:
			children = append(children, t.Content)
		default:
			panic(fmt.Sprintf("assert \"%s\" tag doesn't belong into this namespace HTML", reflect.TypeOf(t).String()))
		}
	}
	return &SvgNode{base.NewFullTagNode(
		"svg",
		model.SVG,
		attributes.Attributes,
		&children,
	)}
}

func AttributesSvg() *svg {
	return &svg{base.NewAttributes()}
}

type svg struct {
	base.Attributes
}

func (s *svg) Width(value string) *svg {
	if value != "" {
		s.RegisterAttribute("width", value, false)
	}
	return s
}
