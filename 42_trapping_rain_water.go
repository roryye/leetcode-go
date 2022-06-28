package main

// https://leetcode.cn/problems/trapping-rain-water/
// 思路1：第 i 个位置蓄水第量 = min(left_max, right_max) - len(i)，时间复杂度为 O(n^2)
// 思路2：假如数组为 [4, 1, 2, 5]，那么 1 和 2 共用左边最高和右边最高的数据，如果数组足够大
// 能复用的概率越大，可维护两个数组，分别为 i 从左（右）找最高的柱子，记录在第 i 的位置，那么
// 第 i 个位置蓄水第量 = min(left[i], right[i]) - len(i) ，时间复杂度为 O（n），空间复杂度为 O（n）
func trap(height []int) int {
	left, right := make([]int, len(height)), make([]int, len(height))
	var leftMax, rightMax int

	for i := 0; i < len(height); i++ {
		leftMax = max(height[i], leftMax)
		left[i] = leftMax
	}
	for i := len(height) - 1; i >= 0; i-- {
		rightMax = max(height[i], rightMax)
		right[i] = rightMax
	}

	var res int
	for i := 0; i < len(height); i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
