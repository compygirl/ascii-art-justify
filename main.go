package main

import (
	"asciirev/asciirev"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]

	var inputStr, data []string // text and ascii data

	banner := ""
	reverse := flag.String("reverse", "", "Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>")
	output := flag.String("output", "", "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	align := flag.String("align", "", "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right something standard")

	flag.Parse()

	if len(input) == 0 {
		fmt.Println("Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>\n\nUsage: go run . [STRING]\nEX: go run . hello")
		return
	}

	if len(input) == 1 {
		// ASCII-ART without BANN ER
		if *reverse == "" {
			inputStr = strings.Split(strings.ReplaceAll(input[0], "\n", "\\n"), "\\n")
			banner = "standard"
			data = asciirev.ReadFile(banner)
			if !asciirev.BannerChecker(banner) {
				return
			}
			fmt.Printf(asciirev.AsciiPrinter(data, inputStr))
		} else {
			// REVERSE
			asciirev.Reverse(*reverse)
		}
	} else if len(input) == 2 {
		// ASCII-ART with BANNER
		inputStr = strings.Split(strings.ReplaceAll(input[0], "\n", "\\n"), "\\n")
		banner := input[1]
		if !asciirev.BannerChecker(banner) {
			return
		}

		data = asciirev.ReadFile(banner)
		fmt.Printf(asciirev.AsciiPrinter(data, inputStr))

	} else if len(input) == 3 {
		if *output != "" {
			inputStr = strings.Split(strings.ReplaceAll(input[1], "\n", "\\n"), "\\n")
			banner := input[2]

			data = asciirev.ReadFile(banner)

			if !strings.HasSuffix(*output, ".txt") {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
				return
			}

			asciirev.Output(*output, inputStr, data)
			return
		} else if *align != "" {
			inputStr = strings.Split(strings.ReplaceAll(input[1], "\n", "\\n"), "\\n")
			banner := input[2]

			data = asciirev.ReadFile(banner)
			asciirev.Align(inputStr, data, *align)
			return
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	}
}
