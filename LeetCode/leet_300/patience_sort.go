package leet_300

import "sort"

/*
	解法2，patience sort，时间复杂度O(NlogN)。
	假设一堆纸牌，A为最大，1为最小。给定一堆纸牌，要求给纸牌排序，规定如下：
		- 小的在上，大的在下。
		- 当新抽出的拍比任何一堆牌的第一张都要小的时候，新建一个堆，把这张牌放到这个位置
		- 重复以上动作，直到用完给定的纸牌为止。

		纸牌的堆数就是最长递增子序列的大小。
 */


func lengthOfLISByPatienceSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var piles [][]int
	for _, num := range nums {
		// Base case, 放第一张牌
		if len(piles) == 0 {
			piles = append(piles, []int{num})
			continue
		}

		// 采用二分查找算法查找新牌应该插入的堆，
		// 该堆堆顶值应该大于等于num
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
