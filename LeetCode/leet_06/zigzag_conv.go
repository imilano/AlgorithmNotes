package leet_06

/*
	题目要将一串字符转换为之字形，转换的字要占用numRows行。

	只需要对s自左向右进行迭代，然后把每个字符添加到合适的行即可。如何确定应该添加到哪一行呢？只需使用两个变量：一个用于标记当前行，一个用于标记当前方向即可。
	当当前行迭代到最底部或者最顶部时，改变方向继续迭代即可。
 */


func convert(s string, numRows int) string {
	// 不加这一句，循环中的if判断就会导致buffer[cur] index out of range.
	if (numRows == 1) {
		return s
	}

	buffer := make([][]int32,numRows)
	cur := 0
	goingDown := false
	for _,c := range s {
		// 将当前字符添加到合适的行
		buffer[cur]  = append(buffer[cur],c)
		// 当进行到最顶部或者最底部时，转变方向
		if cur == 0 || cur == numRows-1 {
			goingDown = !goingDown
		}

		// 递增或递减当前行
		if goingDown {
			cur += 1
		} else {
			cur -= 1
		}
	}

	var result string
	for _,b := range buffer {
		result += string(b)
	}

	return result
}
