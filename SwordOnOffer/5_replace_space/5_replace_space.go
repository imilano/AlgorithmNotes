package __replace_space

// 先在字符串尾部填充任意字符，使得字符串长度等于替换之后的长度。因为一个空格需要替换成三个字符，所以每遇到一个空格，就在末尾添加2个任意字符。
// 令P1指向原字符末尾位置，P2指向填充后的字符串末尾位置，P1和P2向前遍历，当P1指向非空字符时，将P2指向的位置填充为P1指向的值；的那个P1指向空格时，P2依次逆序向前填充02%。
// 当P2追上P1或者P1遍历到字符开始位置时，遍历结束
func ReplaceSpace(s string) string {
	rs := []rune(s)
	n := len(rs)
	oldN := n-1
	for i := n-1; i>= 0; i-- {
		if rs[i] == rune(' ') {
			rs = append(rs,'#','#')
		}
	}

	newN := len(rs)-1
	for oldN >= 0 && oldN < newN {
		c := rs[oldN]
		oldN--
		if c != rune(' ') {
			rs[newN] = c
			newN--
		} else if c == rune(' ') {
			rs[newN] = '0'
			newN--
			rs[newN] = '2'
			newN--
			rs[newN] = '%'
			newN--
		}
	}

	return string(rs)
}

