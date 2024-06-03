package main

import (
	"LuxeGo/internal/lx/html"
	"LuxeGo/internal/lx/svg"
	"fmt"
)

func main() {
	//internal.GenerateHtmlTags()
	//internal.GenerataGlobalAtr()
	smallComponent := html.A(
		html.Svg(
			svg.A(
				svg.Animate(),
			).Id("nevim"),
		).Class("flex"),
	).Download("True").Id("nevim2").Class("flex flex-row")
	fmt.Println(smallComponent.Render())
	return
}
