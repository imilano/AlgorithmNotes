package leet_240

/*
Write an efficient algorithm that searches for a target value in an m x n integer matrix. The matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
*/

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	row, col := rows-1, 0
	for row >= 0 && col < cols {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			col++
		} else if target < matrix[row][col] {
			row--
		}
	}

	return false
}
