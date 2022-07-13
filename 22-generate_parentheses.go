package main

func generateParenthesis(n int) []string {
	var res []string
	gen(&res, "", "(", 0, 0, 2*n)
	return res
}

func gen(res *[]string, str, char string, left, right, all int) {
	if char == "(" {
		left++
	} else {
		right++
	}
	if right > left || left > all/2 {
		return
	}
	str += char
	if len(str) == all {
		if left == right {
			*res = append(*res, str)
		}
		return
	}
	gen(res, str, "(", left, right, all)
	gen(res, str, ")", left, right, all)
	return
}
