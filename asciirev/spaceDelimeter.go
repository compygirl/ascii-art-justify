package asciirev

func SpaceDelimeter(graph []string) []int {
	index := []int{-1}
	count := 0
	intMap := make(map[int][]int, 0)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == ' ' {
				intMap[i] = append(intMap[i], j)
			}
		}
	}

	for i := 0; i < len(intMap[0]); i++ {
		for j := 1; j < len(intMap); j++ {
			for k := 0; k < len(intMap[j]); k++ {
				if intMap[0][i] == intMap[j][k] {
					count++
				}
			}
		}

		if count == len(intMap)-1 {
			index = append(index, intMap[0][i])
		}
		count = 0
	}

	return index
}
