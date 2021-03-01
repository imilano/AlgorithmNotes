package leet_118

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it.
*/

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		pre := len(res[i-1])

		var tmp []int
		for j := 0; j < pre+1; j++ {
			if j == 0 {
				tmp = append(tmp, 1)
			} else if j == pre {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, res[i-1][j]+res[i-1][j-1])
			}
		}

		res = append(res, tmp)
	}

	return res
}
