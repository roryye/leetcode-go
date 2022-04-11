package main

func lengthOfLongestSubstring(s string) int {
	right, res := -1, 0
	mark := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(mark, s[i-1])
		}

		for right + 1 < len(s) && mark[s[right+1]] == 0 {
			mark[s[right+1]]++
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
