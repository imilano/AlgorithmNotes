package leet_36

/*
	Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
		Each row must contain the digits 1-9 without repetition.
		Each column must contain the digits 1-9 without repetition.
		Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

	Note:
		A Sudoku board (partially filled) could be valid but is not necessarily solvable.
		Only the filled cells need to be validated according to the mentioned rules.
*/
func IsValidSudoku(board [][]byte) bool {
	return isValidSudoku(board)
}

// Original wrong answer
// You have misunderstand the problem
// TODO rewrite answer use set data structure
func isValidSudoku(board [][]byte) bool {
	row := len(board)/3
	if row == 0 {
		return false
	}

	for i := 0; i< row; i++ {
		for j := 0; j< row; j++ {
			midX := i*3 + 1
			midY := j*3 + 1

			if !isUnique(board,midX,midY) {
				return false
			}
		}
	}

	return true
}


func isUnique(board [][]byte, x int, y int) bool {
	m := make(map[byte]int)
	for i := x-1;i <= x+1; i++ {
		for j := y-1; j <= y+1; j++ {
			m[board[i][j]]++
		}
	}

	for i,v := range m {
		if v == 9 && i == '.' || v >1 && i != '.' {
			return false
		}

	}

	return true
}