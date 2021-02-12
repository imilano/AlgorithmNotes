package leet_238

/*
Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
*/

// 两趟扫描，分别计算从左到右和从右到左的因子。第一趟先从左到右，然后分别记录每个位置之前的累乘，例如n1[i]存储的是从nums[0]到nums[i-1]之前的乘积；同理对于n2，
// 只不过n2是从右向左扫描的n1，n2[i]存储的是从n2[i+1]到最后一个元素的累乘。最后pro[i]=n1[i]*n2[i].
// time: O(n), space: O(n)
func productExceptSelf(nums []int) []int {
	var res []int
	n1, n2 := []int{1}, []int{1} // 初始化为1，不能为0
	var pro int = 1

	n := len(nums)
	for i := 1; i < n; i++ { // 从左往右计算
		pro *= nums[i-1]
		n1 = append(n1, pro)
	}

	pro = 1
	for i := n - 2; i >= 0; i-- { // 从右向左扫描
		pro *= nums[i+1]
		n2 = append(n2, pro)
	}

	for i := range n1 {
		res = append(res, n1[i]*n2[n-i-1])
	}

	return res
}

// 上述的空间可以进行优化，由于结果最后都要存储在res中，所以我们可以直接将结果扫描到res中，不需要用到两个临时数组
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = 1
	}

	//from left to right
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// from right to left
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
