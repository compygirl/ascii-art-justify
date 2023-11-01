package asciirev

func AsciiMapperRev(mapGraph map[int]string, data []string) string {
	res := ""
	for i := 1; i <= len(mapGraph); i++ {
		for ind, let2 := range data {
			if let2 == mapGraph[i][:len(mapGraph[i])-1] {
				res += string(rune(ind + 32))
			}
		}
	}
	return res
}
