package asciirev

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
	// "golang.org/x/term"
	// tsize "github.com/kopoli/go-terminal-size"
	// "golang.org/x/term"
)

func findMaxWidth(data []string) int {
	max := 0
	for i := 0; i < len(data); i++ {
		if max < len(data[i]) {
			max = len(data[i])
		}
	}
	return max
}

func getTerminalWidth() (int, error) {
	fd := uintptr(syscall.Stdout)

	var size struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&size)),
	)
	if errno != 0 {
		return 0, errno
	}

	return int(size.Col), nil
}

func Align(inputStr, data []string, alignment string) {
	scrWidth, err1 := getTerminalWidth()
	if err1 != nil {
		fmt.Println("Failed to get terminal width:", err1)
		return
	}

	fmt.Println("Terminal width:", scrWidth)
	// scrWidth, _, err := term.GetSize(0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	content := AsciiPrinter(data, inputStr)
	splitted := strings.Split(content, "\n")
	maxTextWidth := findMaxWidth(splitted)
	if scrWidth < maxTextWidth {
		fmt.Println("the text is too Wide for the screen")
		return
	}
	if alignment == "left" {
		fmt.Println(content)
	} else if alignment == "right" {
		res := make([]string, 0)
		for i := 0; i < len(splitted); i++ {
			numOfSpaces := scrWidth - len(splitted[i])
			spaces := strings.Repeat(" ", numOfSpaces)
			res = append(res, spaces+splitted[i])
		}

		for i := 0; i < len(res); i++ {
			fmt.Print(res[i])
		}
	} else if alignment == "center" {
		res := make([]string, 0)
		for i := 0; i < len(splitted); i++ {
			numOfSpaces := scrWidth - len(splitted[i])
			numSpacesLeft, numSpacesRight := 0, 0
			if numOfSpaces%2 == 0 {
				numSpacesLeft, numSpacesRight = numOfSpaces/2, numOfSpaces/2
			} else {
				numSpacesLeft, numSpacesRight = numOfSpaces/2+1, numOfSpaces/2
			}
			spacesL := strings.Repeat(" ", numSpacesLeft)
			spacesR := strings.Repeat(" ", numSpacesRight)
			res = append(res, spacesL+splitted[i]+spacesR)
		}
		for i := 0; i < len(res); i++ {
			fmt.Print(res[i])
		}
	} else if alignment == "justify" {
		for j := 0; j < len(inputStr); j++ {
			words := strings.Fields(inputStr[j])
			arr := []string{inputStr[j]}
			if len(words) > 1 {
				asciiWord := AsciiPrinter(data, arr)
				splittedSub := strings.Split(asciiWord, "\n") // one line out of 8
				numSpacesBetweenWords := (scrWidth - len(splittedSub[1])) / (len(words) - 1)
				ASCIIRES := AsciiPrinter2(data, arr, numSpacesBetweenWords)
				fmt.Printf(ASCIIRES)
			} else {
				fmt.Printf(AsciiPrinter(data, arr))
			}
		}
	}
}

// go get golang.org/x/term
