package leetcode46

import "fmt"

func permute(nums []int) [][]int {
	Size := [...]int{
		0,
		1,
		2,
		3 * 2,
		4 * 3 * 2,
		5 * 4 * 3 * 2,
		6 * 5 * 4 * 3 * 2,
	}
	nlen := len(nums)
	result := make([][]int, 0, Size[nlen])
	combination := make([]int, nlen)
	var backtrack func(last, used, curpos int)
	backtrack = func(last, used, curlen int) {
		// 0是还没开始的
		// curlen 应该是0->nlen
		if curlen > nlen {
			return
		}
		// 遍历当前层的可能
		possibles := nlen - curlen
		p := last

		fmt.Println(">>> possibles", possibles, "pos", p)
		for i := 0; i < possibles; i++ {
			p++
			// 寻找1个未用过的
			for p %= nlen; used&(1<<p) != 0; p %= nlen {
				p++
			}
			combination[curlen] = nums[p]
			fmt.Printf("used:%06b \t%v %d \n", used, combination[0:curlen+1], nums[p])
			if curlen+1 == nlen {
				temp := make([]int, nlen)
				copy(temp, combination)
				result = append(result, temp)
			} else {
				backtrack(p, used|(1<<p), curlen+1)
			}
		}
	}
	backtrack(nlen-1, 0, 0)
	return result

}
