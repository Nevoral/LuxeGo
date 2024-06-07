package main

import (
	"fmt"
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
	"github.com/Nevoral/LuxeGo/internal/LuxeGo/html"
	"github.com/Nevoral/LuxeGo/internal/LuxeGo/svg"
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
	var table []LuxeGo.Content
	for i := 0; i < 8000; i++ {
		var row []LuxeGo.Content
		for j := 0; j < 1000; j++ {
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
		//fmt.Println(val.Render(level))
		_ = val.Render(level)
	}

}
