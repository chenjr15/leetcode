package leetcode200

func numIslands(grid [][]byte) int {
	// dfs 解法，参考 https://leetcode-cn.com/problems/max-area-of-island/
	m := len(grid)
	n := len(grid[0])
	num := 0
	var countArea func(x, y int)
	countArea = func(x, y int) {
		grid[y][x] = '0'
		if (x-1) > -1 && grid[y][x-1] == '1' {
			countArea(x-1, y)
		}
		if (x+1) < n && grid[y][x+1] == '1' {
			countArea(x+1, y)
		}
		if (y-1) > -1 && grid[y-1][x] == '1' {
			countArea(x, y-1)
		}
		if (y+1) < m && grid[y+1][x] == '1' {
			countArea(x, y+1)
		}
		return
	}

	// 逐行遍历找到1
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == '0' {
				continue
			}
			num++
			countArea(x, y)
			// fmt.Println("Area [",x,",",y,"]=",area)
		}
	}
	return num

}
