package asciirev

import (
	"strings"
)

func AsciiPrinter(symbols []string, text []string) string {
	output := ""
	for _, line := range text {
		if isBlankLine(line) {
			output += "\n"
			continue
		}

		for lineIdx := 0; lineIdx < 8; lineIdx++ {
			for _, char := range line {
				output += (strings.Split(symbols[char-32], "\n")[lineIdx])
			}
			output += "\n"
		}
	}
	return output
}

func AsciiPrinter2(symbols []string, text []string, spaces int) string {
	output := ""
	for _, line := range text {
		if isBlankLine(line) {
			output += "\n"
			continue
		}
		flag := false
		for lineIdx := 0; lineIdx < 8; lineIdx++ {

			for _, char := range line {
				if char == 32 && !flag {
					output += strings.Repeat(" ", spaces)
				} else if char != 32 && flag {
					flag = true
				}
				output += (strings.Split(symbols[char-32], "\n")[lineIdx])
			}
			output += "\n"
		}
	}
	return output
}

func isBlankLine(line string) bool {
	return strings.TrimSpace(line) == ""
}
