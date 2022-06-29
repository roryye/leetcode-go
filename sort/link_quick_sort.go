package sort

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// 选出一个值 temp，temp 左边都 < temp，右边都 > temp
// 维护两个指针，一个指着左边区间的最后一个值，一个往右边扫描找比 temp 大的节点
func linkPartition(start, end *Node) *Node {
	i, j := start, start.Next
	temp := start.Val
	for j != end {
		if j.Val < temp {
			i = i.Next
			i.Val, j.Val = j.Val, i.Val
		}
		j = j.Next
	}
	i.Val, start.Val = start.Val, i.Val
	return i
}

func linkQuickSort(start, end *Node) {
	if start != end && start.Next != end {
		mid := linkPartition(start, end)
		linkQuickSort(start, mid)
		linkQuickSort(mid.Next, end)
	}
}

func printLink(n *Node) {
	for n != nil {
		fmt.Print(n.Val, " ")
		n = n.Next
	}
	fmt.Println()
}
