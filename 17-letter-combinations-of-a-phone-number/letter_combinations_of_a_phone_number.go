package leetcode17

func letterCombinations(digits string) []string {
	table := []string{
		// 0
		"",
		// 1
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	// try backtracking
	dlen := len(digits)
	comb := make([]byte, dlen)
	combinations := make([]string, 0)
	var backtrack func(int)
	backtrack = func(deepth int) {
		// deepth->{0,1,...,dlen}
		if deepth == dlen {
			return
		}
		// 当前的字母
		digit := digits[deepth] - '0'

		// 对应每层的所有选择
		digitlen := len(table[digit])
		for i := 0; i < digitlen; i++ {
			// 选择这个数字的每个可能字母
			comb[deepth] = table[digit][i]
			if deepth+1 == dlen {
				combinations = append(combinations, string(comb))
			} else {
				backtrack(deepth + 1)
			}
		}
	}
	backtrack(0)
	return combinations

}
