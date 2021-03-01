package leet_223

/*
Find the total area covered by two rectilinear rectangles in a 2D plane.

Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.

Example:

Input: A = -3, B = 0, C = 3, D = 4, E = 0, F = -1, G = 9, H = 2
Output: 45
Note:

Assume that the total area is never beyond the maximum possible value of int.
*/

// 不重叠的情况只有四种，分别是一个矩形位于另一个矩形的上下左右，这四种情况都可以从坐标的位置做一个区分。其他情况都是两个矩形重叠了，只需要用
//两个矩形的面积和减去重叠矩形的面积，即可得到最终的面积。那么如何得到重叠矩形的面积呢？求重叠矩形的长和宽可以如下，由于交集都是在中间，所以横边的左端点是
//两个矩形左顶点横坐标的较大值，右端点是两个矩形右端点的较小值；同理，竖边的下端点是两个矩形下端点纵坐标的较大值，上端点是两个矩形上顶点纵坐标的较小值。
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum1, sum2 := (C-A)*(D-B), (G-E)*(H-F)
	if C <= E || H <= B || G <= A || F >= D {
		return sum1 + sum2
	}

	return sum1 - (min(C, G)-max(A, E))*(min(D, H)-max(B, F)) + sum2
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
