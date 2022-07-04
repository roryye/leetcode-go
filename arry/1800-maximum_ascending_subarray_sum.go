package arry

func maxAscendingSum(nums []int) int {
	sum, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum += nums[i]
			ans = max(ans, sum)
			continue
		}
		sum = nums[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
