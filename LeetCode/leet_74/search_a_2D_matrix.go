package leet_74

/*
	Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the previous row.
*/

// time complexity: O(max(m,n))
// space complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	row,col := len(matrix),len(matrix[0])
	curRow,curCol := row-1,0

	for curRow >=0 && curCol <= col-1 {
		cur := matrix[curRow][curCol]
		if cur == target {
			return true
		} else if target > cur {
			curCol++
		} else if target < cur {
			curRow--
		}
	}

	return false
}

// as seen from the matrix, the matrix is just a sorted array, so we just need to treat search as binary search
// n * m matrix convert to an array => matrix[x][y] => a[x * m + y]
// an array convert to n * m matrix => a[x] =>matrix[x / m][x % m];
func searchMatrixBinary(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	m,n := len(matrix),len(matrix[0])
	low,high := 0,m*n
	for low < high {
		mid := low + (high-low)/2
		cur := matrix[mid/n][mid%n]

		if  cur == target {
			return true
		} else if cur < target {
			low = mid +1
		} else {
			high = mid
		}
	}

	return false
}