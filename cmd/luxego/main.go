package main

import (
	"LuxeGo/internal/lx"
	"LuxeGo/internal/lx/html"
	"LuxeGo/internal/lx/svg"
	"fmt"
)

func main() {
	Teststructure()
	return
}

func DivByDiv(some string) lx.Content {
	return html.Div(
		html.Div(
			html.Input().
				Class("flex flex-row").
				Id("nevim"),
			html.FreeStr("nevim neco jsem napsal"),
			html.FreeStr(some),
		).Class("flex flex-row"),
	).Data("nevim")
}

func Teststructure() {
	smallComponent := []lx.Content{html.DOCTYPE().Html(),
		html.Html(
			html.Head(),
			html.Body(
				html.A(
					html.Svg(
						svg.A(
							svg.Animate(),
						).Id("nevim"),
					).Class("flex"),
					html.Div(
						DivByDiv("nevim"),
						html.Comment("nevim co za koment napsat"),
					),
				).Download("True").Id("nevim2").Class("flex flex-row"),
			),
		),
	}
	for _, val := range smallComponent {
		level := 0
		fmt.Println(val.Render(level))
	}

}
