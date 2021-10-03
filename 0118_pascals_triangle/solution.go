package solution0118

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	row := []int{1}
	ans[0] = row

	for r := 1; r < numRows; r++ {
		currow := make([]int, r+1)

		prevrow := ans[r-1]
		for col := 0; col < len(currow); col++ {
			up1 := 0
			if col-1 >= 0 {
				up1 = prevrow[col-1]
			}

			up2 := 0
			if col < len(prevrow) {
				up2 = prevrow[col]
			}

			currow[col] = up1 + up2
		}

		ans[r] = currow
	}

	return ans
}
