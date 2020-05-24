package main

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var piles [][]int
	for _, num := range nums {
		// Base case, 初始化堆
		if len(piles) == 0 {
			piles = append(piles, []int{num})
			continue
		}

		// 采用二分查找算法查找新牌应该插入的堆，
		// 该堆堆顶值比num大于等于num
		j := sort.Search(len(piles), func(k int) bool {
			return piles[k][len(piles[k])-1] >= num
		})

		// 如果找到合适的堆，就把牌插入到该堆里;
		// 否则就新建一个堆，并用该值初始化该堆
		if j < len(piles) {
			piles[j] = append(piles[j], num)
		} else {
			piles = append(piles, []int{num})
		}
	}
	// 最后仅需返回piles长度即可
	return len(piles)
}
