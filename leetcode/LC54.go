package leetcode

func spiralOrder(matrix [][]int) []int {
	var res []int
	l := 0
	r := len(matrix[0])
	t := 0
	d := len(matrix)
	for l < r && t < d {
		for i := l; i < r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t >= d {
			break
		}
		for i := t; i < d; i++ {
			res = append(res, matrix[i][r-1])
		}
		r--
		if r <= l {
			break
		}
		for i := r - 1; i >= l; i-- {
			res = append(res, matrix[d-1][i])
		}
		d--
		if d <= t {
			break
		}
		for i := d - 1; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
	}
	return res
}
