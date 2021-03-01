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

	m, n := len(matrix), len(matrix[0])
	length := m * n
	r0, r1 := 0, m-1
	c0, c1 := 0, n-1

	//for r0 <= r1 && c0 <= c1 {
	// use m*n for interview friendly
	for len(res) < length {
		// Add top. 从左走到右
		for top := c0; top <= c1 && len(res) < length; top++ {
			res = append(res, matrix[r0][top])
		}

		// Add right. 从最后一列的第二行走到最后一行
		for right := r0 + 1; right <= r1 && len(res) < length; right++ {
			res = append(res, matrix[right][c1])
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
		for bottom := c1 - 1; bottom >= c0 && len(res) < length; bottom-- {
			res = append(res, matrix[r1][bottom])
		}

		// Add left. 在第一列，从倒数第二行走到正数第二行。
		for left := r1 - 1; left > r0 && len(res) < length; left-- {
			res = append(res, matrix[left][c0])
		}

		r0++
		c0++
		r1--
		c1--
	}

	return res
}


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func printMatrix( matrix [][]int ) []int {
	var res []int
	if matrix == nil {
		return res
	}
	rows := len(matrix)
	var cols int
	if rows != 0 {
		cols = len(matrix[0])
	}

	left,right,top,bottom := 0,cols-1,0,rows-1
	for len(res) < rows * cols {
		for l := left; l <= right;l++ {
			res = append(res, matrix[top][l])
		}

		for r := top + 1; r <= bottom;r++ {
			res = append(res,matrix[r][right])
		}

		if top != bottom {   // 注意：上与下不重合时候才能够从右向左遍历
			for r := right-1; r >= left; r-- {
				res = append(res, matrix[bottom][r])
			}
		}

		if left != right {   // 注意： 左与右不重合时候才能够从下到上遍历
			for b := bottom-1; b > top;b-- {
				res = append(res, matrix[b][left])
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return res
}
