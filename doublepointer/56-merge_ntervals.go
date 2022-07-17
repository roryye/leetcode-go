package doublepointer

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i := 0; i < len(intervals); {
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		res = append(res, []int{intervals[i][0], end})
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
