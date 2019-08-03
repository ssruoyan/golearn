package tool

func Make2DArr(len1 int, len2 int) [][]int {
	m := make([][]int, len1)

	for i := 0; i < len1; i ++ {
		m[i] = make([]int, len2)
	}

	return m
}