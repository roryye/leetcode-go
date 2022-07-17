package tree

func permute(nums []int) [][]int {
	var res [][]int
	var path []int

	// 这里需要指定 sLen，因为在回溯的时候，会对切片 s 进行裁剪操作，导致其长度不稳定
	var dfs func(s []int, sLen int)
	dfs = func(s []int, sLen int) {
		if 0 == len(s) {
			res = append(res, append([]int{}, path...))
		}
		for i := 0; i < sLen; i++ {
			path = append(path, s[i])
			cur := s[i]
			s = append(s[:i], s[i+1:]...)
			dfs(s, len(s))
			s = append(s[:i], append([]int{cur}, s[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, len(nums))
	return res
}
