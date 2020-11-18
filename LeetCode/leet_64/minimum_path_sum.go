package leet_64

/*
	Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
	Note: You can only move either down or right at any point in time.
 */

func min(a,b int) int {
	if a > b {
		return b
	}

	return  a
}

// time complexity: O(m*n)
// space complexity: O(1)
func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	row,col := len(grid),len(grid[0])
	for i := 1; i< row;i++{
		grid[i][0] = grid[i-1][0] +grid[i][0]
	}

	for i := 1; i< col;i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for i := 1; i< row;i++ {
		for j := 1; j< col; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[row-1][col-1]
}