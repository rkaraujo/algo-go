package solution0566

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	i := 0
	j := 0
	for _, row := range mat {
		for _, val := range row {
			result[i][j] = val

			j++
			if j >= c {
				j = 0
				i++
			}
		}
	}

	return result
}
