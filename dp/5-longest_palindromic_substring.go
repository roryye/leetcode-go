package dp

// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	begin, maxLen := 0, 0

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for L := 1; L <= len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := L + i - 1
			if j >= len(s) || j < 0 {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == true && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}
