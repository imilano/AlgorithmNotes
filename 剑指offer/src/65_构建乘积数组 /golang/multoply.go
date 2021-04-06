package golang

// 给定一个数组A[0,1,...,n-1],请构建一个数组B[0,1,...,n-1],其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。
//（注意：规定B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2]; 对于A长度为1的情况，B无意义，故而无法构建，因此该情况不会存在。

//使用两个数组，分别进行左右的扫描，最后再取相关的元素进行相乘即可
// 时间复杂度O(n)，空间复杂度O(n)
func multiply( A []int ) []int {
	// write code here
	n := len(A)
	res := make([]int,n)
	left,right := make([]int, n),make([]int,n)
	left[0] = 1;right[n-1] = 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1]*A[i-1]
	}

	for i := n-2; i >= 0; i-- {
		right[i] = right[i+1]*A[i+1]
	}

	for i := range res {
		res[i] = left[i] * right[i]
	}
	return res
}