package sort

// 因为每次只是收敛一个区间，所以复杂度为
// N + N/2 + N/4 + ... + N/N = N(1-(1/2)^n)/(1-1/2)
// = 2N(1-(1/2)^n) = 2N => O(N)
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
