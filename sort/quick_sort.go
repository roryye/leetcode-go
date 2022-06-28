package sort

// 时间复杂度
// 最优：O(nlogn)
// 最坏：O(n^2)，每次划分只比上次划分找一个数
// 空间复杂度
// 最优：O(logn)
// 最坏：O(n)，
func quickSort(n []int, left, right int) {
	if left < right {
		i := partition(n, left, right)
		quickSort(n, left, i-1)
		quickSort(n, i+1, right)
	}
}

func partition(n []int, left, right int) int {
	temp := n[left]
	for left < right {
		for left < right && n[right] >= temp {
			right--
		}
		n[left] = n[right]

		for left < right && n[left] < temp {
			left++
		}
		n[right] = n[left]
	}
	n[left] = temp

	return left
}
