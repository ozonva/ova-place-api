package utils

func SplitInt(slice []int, batchSize int) [][]int {
	sliceLen := len(slice)
	splittedLen := sliceLen / batchSize
	if sliceLen%batchSize != 0 {
		splittedLen++
	}
	splitted := make([][]int, splittedLen)

	for i := 0; i < splittedLen; i++ {

		if i == 0 {
			splitted[i] = slice[i : i+batchSize]
			continue
		}

		if i*batchSize+batchSize > sliceLen {
			splitted[i] = slice[i*batchSize:]
			continue
		}

		splitted[i] = slice[i*batchSize : i*batchSize+batchSize]
	}

	return splitted
}

func SplitStr(slice []string, batchSize int) [][]string {
	sliceLen := len(slice)
	splittedLen := sliceLen / batchSize
	if sliceLen%batchSize != 0 {
		splittedLen++
	}
	splitted := make([][]string, splittedLen)

	for i := 0; i < splittedLen; i++ {
		if i == 0 {
			splitted[i] = slice[i : i+batchSize]
			continue
		}

		if i*batchSize+batchSize > sliceLen {
			splitted[i] = slice[i*batchSize:]
			continue
		}

		splitted[i] = slice[i*batchSize : i*batchSize+batchSize]
	}

	return splitted
}

func FilterByBlackListInt(slice []int, blackList []int) []int {
	var filtered []int

	for i := range slice {
		if inIntSlice(slice[i], blackList) {
			continue
		}

		filtered = append(filtered, slice[i])
	}

	return filtered
}

func FilterByBlackListStr(slice []string, blackList []string) []string {
	var filtered []string

	for i := range slice {
		if inStrSlice(slice[i], blackList) {
			continue
		}

		filtered = append(filtered, slice[i])
	}

	return filtered
}

func inIntSlice(needle int, haystack []int) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}

	return false
}

func inStrSlice(needle string, haystack []string) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}

	return false
}
