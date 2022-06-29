package string

func lengthOfLongestSubstring(s string) int {
	right, res := -1, 0
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for right+1 < len(s) && m[s[right+1]] == false {
			m[s[right+1]] = true
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
