package golang

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数即 999。

func PrintNumber(n int) []string{
	var res []string
	helper(&res,nil,0,n)
	return res
}

func helper(res *[]string,cur []rune, start int, total int) {
	if start >= total {
		if len(cur) != 0 {
			*res = append(*res, printNum(cur))
		}
		return
	}

	var i rune = '0'
	for ; i <= '9'; i++ {
		helper(res,append(cur,i),start+1,total)
	}
}

func printNum(rs []rune) string {
	var index int
	n := len(rs)
	for index < n && rs[index] == '0' {
		index++
	}

	return string(rs[index:])
}