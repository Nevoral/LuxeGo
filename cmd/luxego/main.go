package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Nevoral/LuxeGo"
	"github.com/Nevoral/LuxeGo/html"
	"github.com/Nevoral/LuxeGo/svg"
	"time"
)

func main() {
	start := time.Now()

	Teststructure()
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	return
}

func DivByDiv(some string) LuxeGo.Content {
	var table []any
	for i := 0; i < 2; i++ {
		var row []any
		for j := 0; j < 2; j++ {
			row = append(row, html.Div(html.FreeStr(fmt.Sprintf("i: %d, j: %d", i, j))))
		}
		table = append(table, row...)
	}
	return html.Div(
		html.Div(
			html.Input().
				Class("flex flex-row").
				Id("nevim"),
			html.FreeStr("nevim neco jsem napsal"),
			html.Div(table...),
			html.FreeStr(some),
		).Class("flex flex-row"),
	).Data("", "nevim")
}

func Teststructure() {
	smallComponent := []LuxeGo.Content{html.DOCTYPE().Html(),
		html.Html(
			html.Head(),
			html.Body(
				html.Comment("nevim nějakydfakusdfhkasdf"),
				html.A(
					html.Svg(
						svg.A(
							svg.Animate(),
						).Id("nevim"),
					).ViewBox("").Fill("").Class("flex"),
					html.Div(
						"nevim",
						"html.Div(),",
						DivByDiv("nevim"),
						//html.Comment("nevim co za koment napsat"),
					),
				).Download("True").Id("nevim2").Class("flex flex-row"),
			),
		),
	}
	var output bytes.Buffer
	for _, val := range smallComponent {
		fmt.Println(val.RenderString(0))
		if err := val.Render(context.Background(), &output); err != nil {
			panic(err)
		}
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println(output.String())

}
