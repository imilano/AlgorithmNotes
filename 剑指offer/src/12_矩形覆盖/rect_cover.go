package _2_矩形覆盖


// 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？比如n=3时，2*3的矩形块有3种覆盖方法.
// 当n为1时，只有一种覆盖方法。
// 当n为2时，有两种覆盖方法。
// 要覆盖2*n的大矩形，可以先覆盖2*1的矩形，再覆盖2*(n-1)的矩形；或者先覆盖2*2的矩形，再覆盖2*(n-2)的矩形。而覆盖2*(n-1)和2*(n-2)的矩形又
// 可以看做子问题。 故而得到类似Fibnacci数列一样的递推公式
func rectCover( number int ) int {
	// write code here
	if number <= 2 {
		return number
	}

	N,NMinusOne,NMinusTwo := 0,2,1
	var count int
	for count < number-2 {
		N = NMinusOne + NMinusTwo
		NMinusTwo = NMinusOne
		NMinusOne = N
		count++
	}

	return N
}