package leet_59

/*
	Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
	Example:
	Input: 3
	Output:
	[
	 [ 1, 2, 3 ],
	 [ 8, 9, 4 ],
	 [ 7, 6, 5 ]
	]
*/

func generateMatrix(n int) [][]int {
	var res [][]int
	var val int
	r0, r1 := 0, n-1
	c0, c1 := 0, n-1
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}

	val = 1
	for val <= n*n && res != nil {
		// from left to right
		for up := c0; up <= c1; up++ {
			res[r0][up] = val
			val++
		}

		// up to bottom
		for right := r0 + 1; right <= r1; right++ {
			res[right][c1] = val
			val++
		}
		// right to left
		for bottom := c1 - 1; bottom >= c0; bottom-- {
			res[r1][bottom] = val
			val++
		}

		// bottom to up
		for left := r1 - 1; left >= r0+1; left-- {
			res[left][c0] = val
			val++
		}

		r0++
		r1--
		c0++
		c1--
	}

	return res
}
