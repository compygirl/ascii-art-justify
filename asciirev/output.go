package asciirev

import (
	"os"
)

func Output(filename string, inputStr, data []string) {
	output := AsciiPrinter(data, inputStr)
	if len(filename) != 0 {
		os.WriteFile(filename, []byte(output), 0o644)
	}
}
