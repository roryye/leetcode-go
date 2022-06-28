package sort

func adjustHeap(a []int, size, n int) {
	leftChild, rightChild := 2*n+1, 2*n+2
	largest := n

	if leftChild < size && a[leftChild] > a[largest] {
		largest = leftChild
	}
	if rightChild < size && a[rightChild] > a[largest] {
		largest = rightChild
	}

	if largest != n {
		a[n], a[largest] = a[largest], a[n]
		adjustHeap(a, size, largest)
	}
}

func buildMaxHeap(a []int) {
	for i := len(a) / 2; i >= 0; i-- {
		adjustHeap(a, len(a), i)
	}
}

func heapSort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, i, 0)
	}
}
