package leet_73

/*
	Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it in-place.
	Follow up:
		A straight forward solution using O(mn) space is probably a bad idea.
		A simple improvement uses O(m + n) space, but still not the best solution.
		Could you devise a constant space solution?
*/

// 使用两个变量代表行和列的状态，如果该行包含0，那就就把行状态设为0，如果该列包含0，那么就把列状态设为0。
// 如果arr[i][j]为0，那么我们就设置行开头和列开头为0，也就是说，设置arr[i][0]为0，设置arr[0][j]为0。对于arr[0][0]而言，因为二者的状态都重叠在arr[0][0]这个
// 格子里，所以我们使用一个变量代表col0，从而做到常量空间的占用
func setZeroes(matrix [][]int)  {
	col0,row,col := 1,len(matrix),len(matrix[0])

	// 第一轮，先探测并且设置每一行和每一列的状态
	for i:= 0; i< row; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j< col;j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := row -1; i>=0;i-- {
		for j := col-1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] =0
			}
		}

		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}