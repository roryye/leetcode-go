package main

func findTopMInN(n []int, m int) int {
	for i := 0; i < len(n); i++ {
		if n[i] == i {
			if (m - 1) == i {
				return n[i]
			}
			continue
		}
		for n[i] != i {
			j := n[i]
			n[i], n[n[i]] = n[n[i]], n[i]
			if (m - 1) == j {
				return n[j]
			}
		}
	}
	return 0
}
