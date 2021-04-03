package _7_数组中的逆序对

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对
// 1000000007取模的结果输出。 即输出P%1000000007



func InversePairs( data []int ) int {
	// write code here
	return brutalForce(data)
}

func brutalForce(data []int) int {
	var res int
	n := len(data)

	for i := 0; i < n-1;i++ {
		for j := i+1; j < n;j++ {
			if data[i] > data[j] {
				res++
			}
		}
	}

	return res
}
