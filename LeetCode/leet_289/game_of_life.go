package leet_289

/*
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
Write a function to compute the next state (after one update) of the board given its current state. The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously.

Example:

Input:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
Output:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
Follow up:

Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?
Credits:
Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
*/

// 最简单的办法就是重新创建一个数组，然后扫描原数组来更新新数组
func gameOfLife(board [][]int) {
	row := len(board)
	var col int
	if row > 0 {
		col = len(board[0])
	}

	dx, dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}, []int{-1, 0, 1, -1, 1, -1, 0, 1}
	matrix := make([][]int, row)
	for i := range board {
		matrix[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			matrix[i][j] = board[i][j]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var cnt int
			for k := 0; k < 8; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < row && y >= 0 && y < col && board[x][y] == 1 {
					cnt++
				}
			}

			if (cnt < 2 || cnt > 3) && board[i][j] == 1 {
				matrix[i][j] = 0
			}

			if (cnt == 2 || cnt == 3) && board[i][j] == 1 {
				matrix[i][j] = 1
			}

			if cnt == 3 && board[i][j] == 0 {
				matrix[i][j] = 1
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = matrix[i][j]
		}
	}
}

// follow-up要求in-place更新，然而在循环中我们肯定是逐个更新的。那么如何记住已经更新过的cell之前的状态呢？这里可以使用状态机：
//- 状态0：死细胞转为死细胞
//- 状态1： 活细胞转为活细胞
//- 状态2： 活细胞转为死细胞
//- 状态3： 死细胞转为活细胞
//
//最后所有状态对2取余，那么状态0和2变为死细胞，状态1和3变为活细胞。首先扫描数组，统计其活细胞邻居的个数，如果遇到状态1和2，则cnt累加。扫描完
//当前cell后，如果活细胞少于2或者大于3，且当前细胞是活细胞的话，则标记状态2；如果当前细胞有三个活细胞邻居且当前细胞恰为死细胞的话，则标记为状
//态3。完成所有扫描后再对数据扫描一遍，对2取余即为我们需要的结果。

// 好巧妙，面试想不出来
func gameOfLife2(board [][]int) {
	row := len(board)
	var col int
	if row != 0 {
		col = len(board[0])
	}

	dx, dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}, []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var cnt int
			for k := 0; k < 8; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < row && y >= 0 && y < col && (board[x][y] == 1 || board[x][y] == 2) {
					cnt++
				}
			}

			if board[i][j] == 1 && (cnt < 2 || cnt > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && cnt == 3 {
				board[i][j] = 3
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] %= 2
		}
	}
}
