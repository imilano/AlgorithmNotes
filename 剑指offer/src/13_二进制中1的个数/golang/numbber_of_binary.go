package golang

// 这题目很简单，简单的进行遍历即可
func NumberOf1( n int ) int {
	// write code here
	var res int

	for i := 0; i < 32; i++ {
		if n & 1 == 1 {
			res++
		}

		n >>= 1
	}

	return res
}

