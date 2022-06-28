package sort

func quickSelect(n []int, left, right, k int) int {
	i, j := left, right
	temp := n[i]
	for i < j {
		for i < j && n[j] >= temp {
			j--
		}
		n[i] = n[j]

		for i < j && n[i] < temp {
			i++
		}
		n[j] = n[i]
	}
	n[i] = temp

	if i == k {
		return n[i]
	} else if i < k {
		return quickSelect(n, i+1, right, k)
	} else {
		return quickSelect(n, left, i-1, k)
	}
}
