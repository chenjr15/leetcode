package leetcode695

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	area := 0
	var countArea func(x, y int)
	countArea = func(x, y int) {
		area++
		grid[y][x] = 0
		if (x-1) > -1 && grid[y][x-1] == 1 {
			countArea(x-1, y)
		}
		if (x+1) < n && grid[y][x+1] == 1 {
			countArea(x+1, y)
		}
		if (y-1) > -1 && grid[y-1][x] == 1 {
			countArea(x, y-1)
		}
		if (y+1) < m && grid[y+1][x] == 1 {
			countArea(x, y+1)
		}
		return
	}

	// 逐行遍历找到1
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 0 {
				continue
			}
			area = 0
			countArea(x, y)
			// fmt.Println("Area [",x,",",y,"]=",area)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea

}
