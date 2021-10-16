package leetcode40

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {

	// 用上一题的直接回溯方法不太奏效，主要是因为候选有重复且结果要去重。
	// 考虑这样的情况candidates = [1,2,1], target = 3, 上一题的方法是直接考虑每个candidate要选或者不选
	// 因此用上一题的方法会出现 [1,2],[2,1]两种结果。直接的解决办法当然是给结果排序然后去重，
	// 但是这样时间和空间复杂度可能会有点高，要对每个结果都排序。
	// 那我们能不能避免对每个输出都排序呢，对输入排序？
	// 首先考虑对candidates排序，得到[1,1,2]，这样的话得到的结果是[1,2],[1,2]显然省掉了一次排序
	// 但是这还不够，因为还要第结果进行去重。如何直接生成唯一的结果？
	// 我们现在对1会做两次判断，每次碰到1都判断一次要不要加，而不知道前面有没有加过1，加过几个1
	// 如果我们把candidates的表示方法改成[1*2,2*1], 即两个1，和1个2这种表示，
	// 那么每次的选择过程就从选或者不选，变成了选0个，选1个和选2个，直接一次把所有的1解决
	// 这样就只会得到[2],[1,2],[1,1,2]三种候选，只有第二种是符合结果的，结果自然不会重复了

	// 这里还是需要对candidates进行一次排序，就算用排序的复杂度也是O(nlogn),
	// 能不能把排序再优化一下？考虑到题目中candidates的返回是给定的50，
	// 我们可以直接建一个长度为50的数组，存放每个数字出现的个数。遍历一次candidates就可以构造好该数组了。
	// 这样只要o(n)就可以吧原来的candidates的表示转换成[1*2,2*1], 即两个1，和1个2这种表示，
	// 这个其实就是计数排序的思想

	clen := len(candidates)
	results := make([][]int, 0)
	combination := make([]int, 0)
	candidateCnt := [51]int{}
	for i := 0; i < clen; i++ {
		candidateCnt[candidates[i]]++
	}
	fmt.Println(candidateCnt)
	var temp []int = nil

	// 回溯法,固定遍历每个数字加或者不加
	var backtrack func(deepth int, curSum int)
	backtrack = func(deepth int, curSum int) {
		if deepth > 50 {
			// 遍历完所有候选
			return
		}
		cnt := candidateCnt[deepth]

		if cnt == 0 {
			backtrack(deepth+1, curSum)
			return
		}

		tempSum := 0
		for i := 1; i <= cnt; i++ {
			tempSum = curSum + i*deepth
			// fmt.Printf("%d个%d,tempsum:%d %v,%d\n",useCnt,deepth,tempSum,combination,cnt)
			if tempSum > target {
				// fmt.Println("break",tempSum,curSum,i)
				// 当前targe加上去都大了，多加几个肯定更大，那就直接break吧
				break
			}

			// 输出
			for k := 0; k < i; k++ {
				combination = append(combination, deepth)
			}
			// fmt.Printf("%d个%d,tempsum:%d %v\n",i,deepth,tempSum,combination)
			if tempSum == target {
				temp = make([]int, len(combination))
				copy(temp, combination)
				results = append(results, temp)
			}
			backtrack(deepth+1, tempSum)
			combination = combination[0 : len(combination)-i]

		}
		backtrack(deepth+1, curSum)

	}
	backtrack(1, 0)
	return results

}
