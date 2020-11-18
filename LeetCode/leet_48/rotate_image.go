package leet_48

/*
	You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
	You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
	DO NOT allocate another 2D matrix and do the rotation.
*/

/* rotate clockwise, first reverse the matrix, then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
 *
 * rotate anti-clockwise
 * first reverse left to right, then swap the symmetry
 * 1 2 3     3 2 1     3 6 9
 * 4 5 6  => 6 5 4  => 2 5 8
 * 7 8 9     9 8 7     1 4 7
 */

func reverse(matrix [][]int) {
	length := len(matrix)
	for i := 0; i< length/2; i++ {
		tmp := matrix[i]
		matrix[i] = matrix[length-i-1]
		matrix[length-i-1] = tmp
	}
}

func swap(matrix [][]int, i,j int) {
	tmp := matrix[i][j]
	matrix[i][j] = matrix[j][i]
	matrix[j][i] = tmp
}

func rotate(matrix [][]int)  {
	reverse(matrix)

	row := len(matrix)
	for i := 0; i< row; i++ {
		for j := i +1; j < row; j++ {
			swap(matrix,i,j)
		}
	}
}

func rotateAntiClock(matrix [][]int) {
	// first , swap leftmost column and rightmost column
	row := len(matrix)
	for i := 0; i< row; i++ {
		tmp := matrix[0][i]
		matrix[i][0] = matrix[i][row-1]
		matrix[i][row-1] = tmp
	}

	// then swap
	for i:= 0; i< row; i++ {
		for j := i +1; j< row; j++ {
			swap(matrix,i,j)
		}
	}
}