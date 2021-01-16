package leet_79

/*
	Given a 2D board and a word, find if the word exists in the grid.
	The word can be constructed from letters of sequentially adjacent cells, where "adjacent" cells are horizontally or vertically neighboring.
	The same letter cell may not be used more than once.
*/

func helper(board [][]byte, word string, i int, j int, index int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 || board[i][j] != word[index] {
		return false
	}
	board[i][j] = '*' // means we have match the byte, this byte has been used , can't be used again
	result := helper(board, word, i+1, j, index+1) ||
		helper(board, word, i-1, j, index+1) ||
		helper(board, word, i, j+1, index+1) ||
		helper(board, word, i, j-1, index+1)
	// 前面的几个递归调用中，我们都修改了board的值，因为修改是会影响到其他的递归子程序的，所以我们需要在退出本递归子程序的时候恢复我们的修改。
	board[i][j] = word[index]
	return result
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}
