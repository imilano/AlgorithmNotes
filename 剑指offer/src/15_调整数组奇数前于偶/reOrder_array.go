package _5_调整数组奇数前于偶

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，
// 偶数和偶数之间的相对位置不变。
func reOrderArray( array []int ) []int {
	// write code here
	var odd,even []int
	for i := range array {
		if array[i] % 2 == 0 {
			even = append(even,array[i])
		} else {
			odd = append(odd, array[i])
		}
	}

	odd = append(odd, even...)
	return odd
}