package leetcode1160

func countCharacters(words []string, chars string) int {

	// 当hashtable用
	charmap := [26]int{}
	m := [26]int{}
	mastered := 0
	for _, c := range []byte(chars) {
		charmap[c-'a'] += 1
	}

	count := 0
	for _, word := range words {
		count = len(word)
		for i := range m {
			m[i] = charmap[i]
		}
		for _, c := range []byte(word) {
			if m[c-'a'] > 0 {
				m[c-'a'] -= 1
			} else {
				count = 0
				break
			}
		}
		mastered += count
	}
	return mastered
}
