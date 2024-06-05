package main

import (
	"LuxeGo/internal"
)

func main() {
	internal.GenerateHtmlTags()
	internal.GenerataGlobalAtr()
	internal.GenerateSvgTags()
	internal.GenerataGlobalAtrSvg()
	return
}

func Teststructure() {
	//smallComponent := html.A(
	//	html.Svg(
	//		html.A(
	//			svg.Animate(),
	//		).Id("nevim"),
	//	).Class("flex"),
	//).Download("True").Id("nevim2").Class("flex flex-row")
	//fmt.Println(smallComponent.Render())
}
