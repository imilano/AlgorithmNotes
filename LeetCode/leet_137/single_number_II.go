package leet_137

/*

Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:
	Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
*/

//使用一个32位的数字来表示每个数位出现1的个数，如果出现的1等于3次，那么对3取余。对每一位都这样操作，最后剩下的就是那个只出现了一次的数字。
// WARNING 不能直接使用int，建议仔细查看go的类型系统
func singleNumber(nums []int) int {
	var res int32
	for i := 0; i <32;i++ {
		var sum int32
		for j := 0; j <len(nums);j++ {
			sum += int32((nums[j] >> i) & 1)  // 求每一位1的个数
		}

		res |= (sum %3) << i
	}

	return int(res)
}
