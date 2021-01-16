package leet_212

/*
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.



Example 1:


Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Example 2:


Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []


Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] is a lowercase English letter.
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
All the strings of words are unique.
*/


// 仍然可以使用字典树来求解。注意是对words构造字典树，而不是对board构造字典树。因为如果对board构造字典树的话，不能实现在board中的向后查找。
// 但是在面试时，不一定能很快写出字典树的求解方法来。此处使用prefix string map的方法，核心思想就是减少查找次数，避免重复查找。

var cords = []struct{
	x int
	y int
} {
	{1,0},
	{-1,0},
	{0,1},
	{0,-1},
}

func findWords(board [][]byte, words []string) []string {
	var cur []byte
	var ans []string

	m := make(map[string]int)
	for _,s := range words {
		for i := 1; i < len(s);i++ {
			m[s[0:i]] = 1
		}
	}

	for _,s := range words {
		m[s] = 2
	}

	row,col := len(board),len(board[0])
	for i := 0; i< row;i++ {
		for j := 0; j < col;j++ {
			search(board,&cur,&ans,m,i,j)
		}
	}

	return ans
}

func search(board [][]byte,cur *[]byte,ans *[]string, m map[string]int,row,col int) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {  // 注意此处的边界限制，是大于等于而不是大于
		return
	}

	c := board[row][col]  // 保存现场
	*cur = append(*cur,c)

	s := string(*cur)
	if _,ok := m[s]; !ok {
		*cur = (*cur)[:len(*cur)-1]
		return
	} else if m[s] == 2 {
		*ans = append(*ans, s)
		m[s] = 1
	}

	board[row][col] = '.' // 修改现场并递归
	for _, cord := range cords {
		search(board,cur,ans,m,row+cord.x,col+cord.y)
	}

	board[row][col] = c  // 恢复现场
	*cur = (*cur)[:len(*cur)-1]
}