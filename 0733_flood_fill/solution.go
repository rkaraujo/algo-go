package solution0733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		visited[i] = make([]bool, len(image[0]))
	}

	return color(image, sr, sc, newColor, image[sr][sc], visited)
}

func color(image [][]int, sr int, sc int, newColor int, oldColor int, visited [][]bool) [][]int {
	if image[sr][sc] == oldColor && !visited[sr][sc] {
		image[sr][sc] = newColor
		visited[sr][sc] = true

		if sr-1 >= 0 {
			image = color(image, sr-1, sc, newColor, oldColor, visited)
		}
		if sr+1 < len(image) {
			image = color(image, sr+1, sc, newColor, oldColor, visited)
		}

		if sc-1 >= 0 {
			image = color(image, sr, sc-1, newColor, oldColor, visited)
		}
		if sc+1 < len(image[0]) {
			image = color(image, sr, sc+1, newColor, oldColor, visited)
		}
	}
	return image
}

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		visited[i] = make([]bool, len(image))
	}

	return color2(image, sr, sc, newColor, visited)
}

func color2(image [][]int, sr int, sc int, newColor int, visited [][]bool) [][]int {
	return nil
}
