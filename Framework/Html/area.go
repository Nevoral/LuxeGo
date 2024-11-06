package html

import (
	"fmt"
	"strings"
	"time"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	assert "github.com/Nevoral/LuxeGo/pkg/Asserts"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

type ShapeEnum int

const (
	Default ShapeEnum = iota
	Circle
	Poly
	Rect
)

func (s *ShapeEnum) String() string {
	return [...]string{"default", "cicle", "poly", "rect"}[*s]
}

type ColorEnum string

const (
	Black ColorEnum = "black"
	Green
	Blue
)

func (c *ColorEnum) String() string {
	return string(*c)
}

func Hex(value string) ColorEnum {
	value = strings.TrimSpace(value)
	if value[0] == '#' {
		value = value[1:]
	}
	if len(value) == 6 {
		var r, g, b uint8
		_, err := fmt.Sscanf(value, "%02x%02x%02x", &r, &g, &b)
		assert.New(err == nil, fmt.Sprintf("Assert invalid HEX color format"))
	}

	if len(value) == 8 {
		var r, g, b, a uint8
		_, err := fmt.Sscanf(value, "%02x%02x%02x%02x", &r, &g, &b, &a)
		assert.New(err == nil, fmt.Sprintf("Assert invalid HEX color format"))
	}
	return ColorEnum(fmt.Sprintf("#%s", value))
}

func RGB(r, g, b uint8) ColorEnum {
	return ColorEnum(fmt.Sprintf("rgb(%d %d %d)", r, g, b))
}

func RGBA(r, g, b, a uint8) ColorEnum {
	assert.New(a < 100, "Assert alpha canal is specified from 0% to 100%.")
	return ColorEnum(fmt.Sprintf("rgb(%d %d %d / %.1f)", r, g, b, float64(a/100)))
}

func AttributesArea() *area {
	return &area{base.NewAttributes()}
}

type area struct {
	base.Attributes
}

func (area *area) Datetime(value time.Time) *area {
	area.RegisterAttribute("datetime", value.String(), false)
	return area
}

func (area *area) Class(value string) *area {
	area.RegisterAttribute("class", value, false)
	return area
}

func (a *area) Id(value ColorEnum) *area {
	a.RegisterAttribute("id", value.String(), false)
	return a
}

func (a *area) Shape(value ShapeEnum) *area {
	a.RegisterAttribute("shape", value.String(), false)
	return a
}

func Area(attributes *area) base.Content {
	return base.NewSelfClosingNode(
		"area",
		model.HTML,
		attributes.Attributes,
	)
}
