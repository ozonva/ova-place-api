package utils

func InvertMappingInt(toInvert map[int]int) map[int]int {
	inverted := make(map[int]int, len(toInvert))
	for i := range toInvert {
		if _, ok := inverted[toInvert[i]]; ok {
			panic("Key already exists")
		}
		inverted[toInvert[i]] = i
	}

	return inverted
}

func InvertMappingStr(toInvert map[string]string) map[string]string {
	inverted := make(map[string]string, len(toInvert))
	for i := range toInvert {
		if _, ok := inverted[toInvert[i]]; ok {
			panic("Key already exists")
		}
		inverted[toInvert[i]] = i
	}

	return inverted
}
