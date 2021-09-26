package leetcode447

func numberOfBoomerangs(points [][]int) int {

	// 就是穷举了
	total := 0
	d := 0
	for i := range points {
		distance := make(map[int]int)
		// 计算p1 到p2的距离, 使用 k,point:=range points 比 直接用数组索引慢
		for k := range points {
			if i == k {
				// 跳过自己到自己的距离
				continue
			}
			// 计算距离
			d = dis(points[i], points[k])
			distance[d]++
		}

		// 计算所有相同距离的边
		for _, cnt := range distance {
			if cnt >= 2 {
				total += cnt * (cnt - 1)
			}
		}
	}

	return total

}

func dis(a, b []int) int {
	// dx^2 + dy^2
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
