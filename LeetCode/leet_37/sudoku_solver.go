package leet_37

/*
	Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

    Each of the digits 1-9 must occur exactly once in each row.
    Each of the digits 1-9 must occur exactly once in each column.
    Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

The '.' character indicates empty cells.
 */

func helper(board [][]byte,row int, col int) bool {
	if  row == 9 {
		return true
	}

	// check next row
	if col >= 9 {
		return helper(board,row+1,0)
	}

	// current position has number, check next column of current row
	if board[row][col] != '.' {
		return helper(board,row,col+1)
	}

	// enumerate every possible choice
	for c := byte('1'); c <=byte('9');c++ {
		// If place c in current position, is it valid?
		// If not valid, then try another character
		if !isValid(board,row,col,c) {
			continue
		}

		// If it is valid ,put it here
		board[row][col] = c

		// Continue to check column under the current choice, if that choice is valid, return true
		if helper(board,row,col+1) {
			return true
		}

		// If it is not valid, then reset it to default, try next schema
		board[row][col] = '.'
	}

	// after all the choice and there is no valid schema, then just return false
	return false
}

func isValid(board [][]byte, row int, col int,val byte) bool {
	for i := 0; i< 9;i++ {
		if board[row][i] == val  || board[i][col] == val {
			return false
		}
	}

	curRow,curCol := row - row%3,col - col%3
	for i := 0; i< 3; i++  {
		for j := 0; j<3; j++ {
			if board[curRow + i] [curCol + j] == val {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte)  {
	helper(board,0,0)
}