package asciirev

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileReader(filename string) ([]string, error) {
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}
	var data []string

	file, err := os.Open(filename)
	if err != nil {
		return data, fmt.Errorf("open file err: %v", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	defer file.Close()

	return data, nil
}

func CheckFormat(graph []string) string {
	for _, str := range graph {
		for ind, _ := range str {
			if str[ind] != 32 && (str[ind] == 'O' || str[ind] == 'o' || str[ind] == '0') {
				return "thinkertoy"
			}
		}
	}

	for _, str := range graph {
		for ind, _ := range str {
			if str[ind] != 32 && (str[ind] == '\\' || str[ind] == '/' || str[ind] == '(' || str[ind] == ')' || str[ind] == 'V') {
				return "standard"
			}
		}
	}
	return "shadow"
}

func Reverse(filename string) {
	graph, err2 := FileReader(filename) // open txt file
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	banner := CheckFormat(graph) // get type of banner
	data := ReadFile(banner)
	GraphicMapper(graph, data)
}
