package leet_229

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Follow-up: Could you solve the problem in linear time and in O(1) space?



Example 1:

Input: nums = [3,2,3]
Output: [3]
Example 2:

Input: nums = [1]
Output: [1]
Example 3:

Input: nums = [1,2]
Output: [1,2]
*/

// 使用map
func majorityElement(nums []int) []int {
	var res []int
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	n := len(nums)
	for k, v := range m {
		if v > n/3 {
			res = append(res, k)
		}
	}

	return res
}

// 摩尔投票法。根据提示，出现次数大于n/3的数字的个数只会小于等于2，绝对不会超过2.摩尔技术法的步骤是，先假设第一个元素是主元，然后对count加1，
// 然后继续遍历下一个元素，如果下一个元素跟前一个元素相同，那么count继续加1，如果不同，则count减1，如果当前count为0，那么将当前元素设置为新的主元，
// 然后继续向后进行遍历。因为当count为0时候，说明和这个元素不同的元素的个数已经跟这个元素的个数相同，这个元素作为主元的可能性已经很小了，故而需要将新的元素设置为主元
// 具体做法是，对数组进行两趟遍历，第一趟遍历找出两个候选数；第二趟重新投票验证这两个候选数是否为符合题意的数。
func majorityElement2(nums []int) []int {
	var res []int
	var a, b, cnt1, cnt2 int

	// 先找出两个候选数
	for _, num := range nums {
		if num == a {
			cnt1++
		} else if num == b {
			cnt2++
		} else if cnt1 == 0 {
			cnt1 = 1
			a = num
		} else if cnt2 == 0 {
			cnt2 = 1
			b = num
		} else {
			cnt1--
			cnt2--
		}
	}

	// 验证两个候选数是否是真正的主元
	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if num == a {
			cnt1++
		} else if num == b {
			cnt2++
		}
	}

	if cnt1 > len(nums)/3 {
		res = append(res, a)
	}

	if cnt2 > len(nums)/3 {
		res = append(res, b)
	}

	return res
}
