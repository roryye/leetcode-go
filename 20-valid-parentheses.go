package main

import "container/list"

var m = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := list.New()
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.PushBack(s[i])
		} else {
			if stack.Len() == 0 {
				return false
			}

			ele := stack.Back()
			char := ele.Value.(byte)
			if s[i] == m[char] {
				stack.Remove(ele)
			} else {
				return false
			}
		}
	}
	if stack.Len() == 0 {
		return true
	}
	return false
}
