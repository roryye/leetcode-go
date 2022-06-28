package sort

func mergeSort(n []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(n, left, mid-1)
	mergeSort(n, mid+1, right)
	merge(n, left, mid, right)
}

func merge(n []int, left, mid, right int) {
	var temp []int
	i, j := left, mid+1
	for i <= mid && j <= right {
		if n[i] <= n[j] {
			temp = append(temp, n[i])
			i++
		} else {
			temp = append(temp, n[j])
			j++
		}
	}
	// 假如左半边还有剩余没排完
	for i <= mid {
		temp = append(temp, n[i])
		i++
	}
	// 假如右半边还有剩余没排完
	for j <= right {
		temp = append(temp, n[j])
		j++
	}
	for i, v := range temp {
		n[left+i] = v
	}
}
