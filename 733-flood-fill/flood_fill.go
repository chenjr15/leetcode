package leetcode733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	var fill func(x, y int)
	fill = func(x, y int) {
		image[x][y] = newColor
		if (x-1) > -1 && image[x-1][y] == oldColor {
			fill(x-1, y)
		}
		if (x+1) < m && image[x+1][y] == oldColor {
			fill(x+1, y)
		}
		if (y-1) > -1 && image[x][y-1] == oldColor {
			fill(x, y-1)
		}
		if (y+1) < n && image[x][y+1] == oldColor {
			fill(x, y+1)
		}
		return
	}
	fill(sr, sc)
	return image
}
