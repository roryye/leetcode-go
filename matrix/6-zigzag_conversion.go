package matrix

import "bytes"

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n, r := len(s), numRows
	t := 2*r - 2
	var x int
	mat := make([][]byte, n)
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n, r := len(s), numRows
	t := 2*r - 2
	c := len(s)/2 + 1
	mat := make([][]byte, n)
	for i := range mat {
		mat[i] = make([]byte, c)
	}

	var x, y int
	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < r-1 {
			x++
		} else {
			x--
			y++
		}
	}

	res := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				res = append(res, ch)
			}
		}
	}

	return string(res)
}
