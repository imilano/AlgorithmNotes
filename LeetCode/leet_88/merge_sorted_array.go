package leet_88

/*
	Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
	Note:
		The number of elements initialized in nums1 and nums2 are m and n respectively.
		You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if nums1 == nil || nums2 == nil {
		return
	}

	len1,len2 := len(nums1),len(nums2)
	total := len1+len2
	for len1 >0 && len2 > 0 && total > 0 {
		if nums1[len1-1] > nums2[len2-1] {
			nums1[total-1] = nums1[len1-1]
			total--
			len1--
		}  else {
			nums1[total-1] = nums2[len2-1]
			len2--
			total--
		}
	}

	// 无需再计算len1
	//for len1 > 0{
	//	nums1[total-1] = nums1[len1-1]
	//	total--
	//	len1--
	//}

	for len2 >= 0 {
		nums1[total] = nums2[len2]
		total--
		len2--
	}
}