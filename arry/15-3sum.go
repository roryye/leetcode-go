package arry

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 0 || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := range nums {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				fmt.Println(nums[i], nums[left], nums[right])
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
