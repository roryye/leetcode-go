package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, tail *ListNode
	var value, carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		value, carry = sum%10, sum/10

		if res == nil {
			res = &ListNode{Val: value}
			tail = res
		} else {
			tail.Next = &ListNode{Val: value}
			tail = tail.Next
		}
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return res
}
