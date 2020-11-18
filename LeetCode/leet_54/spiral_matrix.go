package leet_54

/*
	Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
*/

// time complexity: O(N)
// space complexity: O(N)
// 看做一个个矩形，只需要把矩形的边界元素慢慢相加就可以得到结果
// 首先从左走到右；然后在最后一列，从下一行走到最后一行；然后从前一列开始，在最后一行从倒数第二列走到第一列；然后在第一列，从倒数第二行到正数第二行。
// 注意，以上顺序不可变，否则会出现重复
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	m,n := len(matrix),len(matrix[0])
	length := m *n
	r0,r1 :=0,m-1
	c0,c1 := 0,n-1

	//for r0 <= r1 && c0 <= c1 {
	// use m*n for interview friendly
	for len(res) < length {
		// Add top. 从左走到右
		for top := c0; top <= c1 && len(res) < length ;top++ {
			res = append(res,matrix[r0][top])
		}

		// Add right. 从最后一列的第二行走到最后一行
		for right := r0 +1; right <= r1 && len(res) < length; right++ {
			res = append(res,matrix[right][c1])
		}

		// If we are not left only two row or only two col
		//if r0 < r1 && c0 < c1 {
		//	// Add bottom
		//	for bottom := c1 -1; bottom >= c0;bottom-- {
		//		res = append(res,matrix[r1][bottom])
		//	}
		//
		//	// Add left
		//	for left := r1 -1; left > r0;left-- {
		//		res = append(res, matrix[left][c0])
		//	}
		//}
		// Add bottom. 在最后一行从倒数第二列走到第一列
		for bottom := c1-1; bottom >= c0 && len(res) < length; bottom-- {
			res = append(res, matrix[r1][bottom])
		}

		// Add left. 在第一列，从倒数第二行走到正数第二行。
		for left := r1-1;left > r0 && len(res) < length; left-- {
			res = append(res, matrix[left][c0])
		}


		r0++
		c0++
		r1--
		c1--
	}

	return res
}