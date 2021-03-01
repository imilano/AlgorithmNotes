package leet_130

/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/

//也就是说，将被X包围的O替换为X，但是需要注意的是，边界的O是不被X包含的。
//解决思路：扫描矩阵的四条边，如果扫描到了O，那么从该处开始DFS，将每个O都设置为$，四条边扫描结束后，重新扫描整个矩阵，将扫描到的O都替换为X，然后将扫描到的$都替换为O。
func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])

	// scan to mark
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (i == 0 || i == row-1 || j == 0 || j == col-1) && board[i][j] == 'O' {
				scanAndMark(board, i, j)
			}
		}
	}

	// scan to restore
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '$' {
				board[i][j] = 'O'
			}
		}
	}

	return
}

// dfs, scan and mark
func scanAndMark(board [][]byte, i, j int) {
	row, col := len(board), len(board[0])
	if board[i][j] == 'O' {
		// mark
		board[i][j] = '$'

		// DFS, 注意的是，标记的范围应该是可以包括边界的，也就是说，如果扫描到了边界上的元素也在未闭合的区域，那么该元素也要被标记。
		if i > 0 && board[i-1][j] == 'O' { // i > 0 或者 i < row-1，意味着可以标记边界上的元素；同理对j
			scanAndMark(board, i-1, j)
		}
		if i < row-1 && board[i+1][j] == 'O' {
			scanAndMark(board, i+1, j)
		}
		if j > 0 && board[i][j-1] == 'O' {
			scanAndMark(board, i, j-1)
		}
		if j < col-1 && board[i][j+1] == 'O' {
			scanAndMark(board, i, j+1)
		}
	}
}
