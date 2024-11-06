package main

import (
	"context"
	"fmt"
	"strings"

	. "github.com/Nevoral/LuxeGo/Framework/Html"
	svg "github.com/Nevoral/LuxeGo/Framework/Svg"
	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
)

func HxTarget(value string) (string, string, bool) {
	return "hx-target", value, false
}

func main() {
	var out strings.Builder

	A(
		AttributesOfA().
			Class("some").
			Id("1").
			CustomAttribute(HxTarget("something")),
		Area(AttributesArea().
			Id(RGB(200, 30, 255)).
			Shape(Circle),
		),
		A(AttributesOfA().
			Class("second"),
			"tags ...any",
			Comment("something"),
			svg.Svg(svg.AttributesSvg().
				Width("5"),
				svg.Svg(svg.AttributesSvg()),
				Comment("CommentS"),
				"askjdfakdsf",
			),
		),
	).Render(context.Background(), &out, base.DefaultHtml)

	fmt.Println(out.String())
}
