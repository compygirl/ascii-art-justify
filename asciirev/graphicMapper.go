package asciirev

import (
	"fmt"
	"log"
)

func GraphicMapper(graph []string, data []string) {
	answer := ""
	for len(graph) > 0 {
		for len(graph) != 0 && graph[0] == "" {
			answer += "\n"
			graph = graph[1:]
		}
		tempGraph := graph[:8]

		index := SpaceDelimeter(tempGraph)
		mapGraph := MapSlicer(tempGraph, index)
		answer += AsciiMapperRev(mapGraph, data)

		graph = graph[8:]
		if len(graph) == 0 {
			break
		} else {
			answer += "\n"
		}
		for len(graph) != 0 && graph[0] == "" {
			answer += "\n"
			graph = graph[1:]
		}

	}
	answer += "\n"
	fmt.Print(answer)
}

func MapSlicer(graph []string, index []int) map[int]string {
	mapGraph := map[int]string{}
	line := ""
	ind := 1

	start := 0
	for i := 0; i < len(index)-1; i++ {
		if index[i]+1 == index[i+1] {
			start = i + 1
		}
		if start != 0 {
			index = append(index[0:start], index[start+5:]...)
			start = 0

		}
	}

	for len(index) != 1 {
		for i := 0; i < len(graph); i++ {
			if len(graph[i]) == 0 {
				error := "Error: please check if the symbols are correctly formated"
				fmt.Println("Note: each symbol in test file should have 8 lines and each line have spaces even though symbol is not long enough")
				log.Fatal(error)
			}
			line += graph[i][index[0]+1 : index[1]+1]
			line += "\n"

		}
		mapGraph[ind] = line
		line = ""
		index = index[1:]

		ind++

	}
	return mapGraph
}
