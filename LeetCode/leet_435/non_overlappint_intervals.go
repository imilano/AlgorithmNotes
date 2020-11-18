package leet_435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// 按照区间的end进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 此处不判断长度的话，对intervals[0][1]进行访问就会导致访问越界
	length := len(intervals)
	if length == 0 {
		return 0
	}

	// 至少有一个区间不相交
	count,end,start:= 1,intervals[0][1],0

	for i := 0; i< length; i++ {
		// 找到下一个区间
		start = intervals[i][0]
		if start >= end {
			count++
			end = intervals[i][1]
		}
	}

	return length - count
}