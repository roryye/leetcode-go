package tree

func combinationSum(candidates []int, target int) [][]int {
	var comb []int
	var res [][]int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}
		// 跳过当前 index
		dfs(target, index+1)

		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}
