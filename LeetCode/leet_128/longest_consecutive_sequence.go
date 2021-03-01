package leet_128

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
Follow up: Could you implement the O(n) solution?
*/

// sequence mean that the number need not to be continuous

func longestConsecutive(nums []int) int {
	return lcsWithSet(nums)
}

// Time complexity : O(n)
// 仔细看下面的第二个for循环，进入第二个for循环其实是有条件的，只有当k-1不出现在集合中时才会进入，仔细观察会发现，每个元素都只被遍历过一次
// 故而时间复杂度为O(n)
func lcsWithSet(nums []int) int {
	// 将数组去重，装进集合中，这样在后续的查找中就能实用O(1)的时间复杂度
	set := make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}

	// 对于每一个在集合中的元素k，如果k-1不在集合中，那么我们就可以用k作为起点，来查找从k开始的最长子序列，直到出现非连续元素g，
	// 那么g-k即为本次查找所得到的最长连续子序列
	var best int
	for k := range set {
		if _, ok := set[k-1]; !ok {
			// 以k为起点，查找k+1，k+2，k+3...k+n是否在集合中
			y := k + 1
			_, exist := set[y]
			for exist {
				y += 1
				_, exist = set[y]
			}

			// 更新最长子序列的长度
			best = max(best, y-k)

		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// wrong
//func lcsWithSort(nums []int) int {
//	res := 1
//	sort.Ints(nums)
//
//	var slow int
//	fast := 1
//	for fast < len(nums) {
//		if nums[fast]  == nums[fast-1]+1 {
//			res = max(res,fast-slow+1)
//			fast++
//		} else {
//			slow = fast
//			fast++
//		}
//	}
//
//	res = max(res, fast -slow )
//
//	return res
//}
//
