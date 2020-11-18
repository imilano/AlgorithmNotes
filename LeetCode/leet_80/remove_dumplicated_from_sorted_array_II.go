package leet_80

/*
	Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
	Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.
 */

//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	left,right,length := 0,0,len(nums)
//	var count int
//	for right < length-1 {
//		if nums[right]  == math.MinInt64 {
//			break
//		}
//		if nums[left] == nums[right] && right - left <= 1 {
//			right++
//		} else if nums[left] != nums[right] && right - left <= 2{
//			left = right
//			right++
//		} else if nums[left] == nums[right] && right - left > 1{
//			left = right
//			for i := right; i < length;i++ {
//				if nums[i] != nums[right] {
//					right = i
//					break
//				}
//			}
//			curCount := right - left
//			for i := left; i< length;i++ {
//				nums[i] = nums[right]
//				right++
//			}
//
//			for i := length-1; i> length-1-curCount;i--{
//				nums[curCount] = math.MinInt64
//			}
//
//			count += curCount
//		}
//	}
//
//	return length-count
//}

// i 和 v开始时指向同一个位置，当当前v的值不比他前第k个数大时（此时i指向v），说明出现了k个及其以上的重复，此时i不变，等找到比他大的数字来替换；当找到比i大的数字时，替换i，然后i向前走；
// 此时循环的 v > nums[i-2] 继续用于保证当前不会出现三个及其以上的重复
func removeDuplicates(nums []int) int {
	var i int
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}

	return i
}

// 提取模板，当最多允许k个重复时
func atMostKDup(nums []int, k int) {
	var i int
	for _, n := range nums {
		if i < k ||  n > nums[i-k] {
			nums[i] = n
			i++
		}
	}
}