package dp

func canJump(nums []int) bool {
	var rightMost int
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = max(rightMost, nums[i]+i)
		}
		if rightMost >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
