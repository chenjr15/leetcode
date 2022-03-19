package leetcode1629

func slowestKey(releaseTimes []int, keysPressed string) byte {
	pre_t := 0
	key_duration := [26]int{0}
	var max_i byte = 0
	duration := 0
	for i, t := range releaseTimes {
		duration = t - pre_t
		pre_t = t
		// fmt.Println(duration)
		if duration > key_duration[max_i] || ((keysPressed[i]-'a') > max_i && duration == key_duration[max_i]) {
			max_i = keysPressed[i] - 'a'
			key_duration[max_i] = duration
		}
	}
	return max_i + 'a'

}
