package main

import (
	"bufio"
	"fmt"
	"github.com/Nevoral/LuxeGo/internal"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter tags name that you want to generate.\nPrompt should looks like this: \"html/svg, svg/a, ...\"\nIf you want all the tags just press enter. Capishto:\n")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim the newline character from the input
	if len(input) > 2 {
		input = input[:len(input)-1]
		tags := strings.Split(input, ", ")

		for _, tag := range tags {
			values := strings.Split(tag, "/")
			if values[0] == "html" {
				internal.GenerateHtmlTag(values[1])
			} else {
				internal.GenerateSvgTag(values[1])
			}
		}
		return
	}
	internal.GenerateHtmlTags()
	internal.GenerataGlobalAtr()
	internal.GenerateSvgTags()
	internal.GenerataGlobalAtrSvg()
	return
}
