package leet_63

/*
	A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
	The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner
	of the grid (marked 'Finish' in the diagram below).
	Now consider if some obstacles are added to the grids. How many unique paths would there be?
	An obstacle and empty space is marked as 1 and 0 respectively in the grid.
	Note: m and n will be at most 100.
 */



/*
	A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
	The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner
	of the grid (marked 'Finish' in the diagram below).
	Now consider if some obstacles are added to the grids. How many unique paths would there be?
	An obstacle and empty space is marked as 1 and 0 respectively in the grid.
	Note: m and n will be at most 100.
*/

// space complexity: O(m*n)
// time complexity: O(1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row,col := len(obstacleGrid),len(obstacleGrid[0])
	// if the starting point has obstacle, then we can't walk to anywhere
	if obstacleGrid[0][0] == 1{
		return 0
	}

	// if the starting point has no obstacle, then the number of path walking to starting point is 1
	obstacleGrid[0][0] =1

	// for first column, if the cell has no obstacle , then set the number of path of that cell to 1, otherwise 0
	for i := 1; i< row; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1{
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	// for the first row, if the cell has no obstacle, then set the number of path of that cell to 1, otherwise 0
	for i := 1; i< col; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1{
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	// every cell can only be reached from up or left, if the cell has obstacle, we set the number of path walking to that cell to 0, otherwise
	// set it to the sum of up and left cell
	for i := 1; i< row; i++ {
		for j := 1; j < col;j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[row-1][col-1]
}
