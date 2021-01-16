package leet_51

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.
Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
*/

//需要注意的是，皇后在的那一行，那一列以及两个对角线上都不允许有皇后。
//因为我们是从第一行开始进行放置的，当我们检测第i行时，我们没有必要检测i行之后的行，我们只需要检测i行之前的行就可以了
// 以行为单位进行检测，在本行放置了皇后之后，本行已经不能再放置皇后了，直接跳转到下一行进行检测。
// 如果在本行的当前列放置了皇后，但是检测时发生了冲突，那么就在当前位置前进一格，然后继续进行检测。
func solve(chess [][]string, res *[][]string, row int) {
	if row == len(chess) {
		s := construct(chess)
		*res = append(*res, s)
		return // don't forget to return the function
	}

	// 从row开始，是因为我们需要的是所有的解决方案，因而意味着本行的每一列都需要进行检测
	// 因为是采用回溯的思想，我们先假设ches[i][j]会放置一个皇后，然后对i，j进行检测，看看是否会冲突。如果会冲突，那么我们跳过当前位置，
	// 直接进入下一格，也就是j+1，然后继对当前的i和j进行检测；如果不冲突，那么我们就将当前的chess[i][j]设置为Q，然后检测下一行，继续递归；
	// 在递归返回后，我们已经知道了如果放置在当前位置的结果，那么我们需要恢复当前位置为.，然后继续进行下一格的匹配
	// The essence of backtrace
	for col := 0; col < len(chess); col++ {
		if isValid(chess, row, col) {
			chess[row][col] = "Q"
			solve(chess, res, row+1)
			chess[row][col] = "."
		}
	}
}

func construct(chess [][]string) []string {
	var path []string
	for i := 0; i < len(chess); i++ {
		var s string
		for j := range chess[i] {
			s += chess[i][j]
		}

		path = append(path, s)
	}

	return path
}
func isValid(chess [][]string, row int, col int) bool {
	// check col
	for i := 0; i < row; i++ {
		if chess[i][col] == "Q" {
			return false
		}
	}

	// check 45°
	for i, j := row-1, col+1; i >= 0 && j < len(chess); {
		if chess[i][j] == "Q" {
			return false
		}
		i--
		j++
	}

	// check 135°
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if chess[i][j] == "Q" {
			return false
		}

		i--
		j--
	}

	return true
}

func solveNQueens(n int) [][]string {
	// initialize chess all to .
	chess := make([][]string, n)
	for i := range chess {
		chess[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chess[i][j] = "."
		}
	}

	var res [][]string
	solve(chess, &res, 0)
	return res
}
