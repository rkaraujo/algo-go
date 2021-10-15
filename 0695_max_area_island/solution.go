package solution0695

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	visited := createVisited(len(grid), len(grid[0]))

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			curArea := area(grid, x, y, visited)
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}

	return maxArea
}

func area(grid [][]int, x int, y int, visited [][]bool) int {
	curArea := 0
	if grid[x][y] == 1 && !visited[x][y] {
		curArea = 1
		visited[x][y] = true

		if x-1 >= 0 {
			curArea += area(grid, x-1, y, visited)
		}
		if x+1 < len(grid) {
			curArea += area(grid, x+1, y, visited)
		}

		if y-1 >= 0 {
			curArea += area(grid, x, y-1, visited)
		}
		if y+1 < len(grid[0]) {
			curArea += area(grid, x, y+1, visited)
		}
	}
	return curArea
}

func createVisited(m, n int) [][]bool {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return visited
}
