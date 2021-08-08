package utils

func InvertMappingInt(toInvert map[int]int) map[int]int {
	inverted := map[int]int{}
	for i := range toInvert {
		if _, ok := inverted[toInvert[i]]; ok {
			panic("Key already exist")
		}
		inverted[toInvert[i]] = i
	}

	return inverted
}

func InvertMappingStr(toInvert map[string]string) map[string]string {
	inverted := map[string]string{}
	for i := range toInvert {
		if _, ok := inverted[toInvert[i]]; ok {
			panic("Key already exist")
		}
		inverted[toInvert[i]] = i
	}

	return inverted
}
