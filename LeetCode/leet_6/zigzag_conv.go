package leet_06

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
*/

//题目要将一串字符转换为之字形，转换的字要占用numRows行。
//只需要对s自左向右进行迭代，然后把每个字符添加到合适的行即可。如何确定应该添加到哪一行呢？只需使用两个变量：一个用于标记当前行，一个用于标记当前方向即可。
//当当前行迭代到最底部或者最顶部时，改变方向继续迭代即可。

func convert(s string, numRows int) string {
	var res string
	if numRows == 1 { // 不加这一句，循环中的if判断就会导致buffer[cur] index out of range
		return s
	}

	buffer := make([][]int32, numRows)
	cur, down := 0, false // cur 标记当前行，down标记本次扫描是否向下走，初始时down标记为false
	for _, c := range s {
		buffer[cur] = append(buffer[cur], c) // 将当前字符添加到合适的行
		if cur == 0 || cur == numRows-1 {    // 当进行到最顶部或者最底部时，转变方向
			down = !down
		}

		if down { // 如果正在向下走，则递增行数，负责递减行数
			cur += 1
		} else {
			cur -= 1
		}
	}

	for _, b := range buffer {
		res += string(b)
	}

	return res
}
