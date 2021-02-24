package leet_294

/*
You are playing the following Flip Game with your friend: Given a string that contains only these two characters: + and -, you and your friend take turns to flip two consecutive "++" into "--". The game ends when a person can no longer make a move and therefore the other person will be the winner.

Write a function to determine if the starting player can guarantee a win.

Example:

Input: s = "++++"
Output: true
Explanation: The starting player can guarantee a win by flipping the middle "++" to become "+--+".
Follow up:
Derive your algorithm's runtime complexity.
*/

// 回溯法。检查是否存在一种选择，能够让p1 win，但是无论p2怎么选都不能win；如果当前情况下，无论p1怎么选，p2都能win，那么p1肯定输
func canWin(s string) bool {
	for i := 1; i< len(s);i++ {
		if s[i] == '+' && s[i-1] == '+' && !canWin(s[:i-1]+"+"+"+"+s[i+1:]) {
			return true  // 在当前的情况下，是否存在一种选择，使得p1翻转之后，无论p2怎么翻转，都不能win
		}
	}

	return false  //在当前的情况下，无论p1怎么翻转，p1都不可能win
}