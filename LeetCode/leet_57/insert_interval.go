package leet_57

import (
	"sort"
)

/*
	Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
	You may assume that the intervals were initially sorted according to their start times.
 */

// original solution
// time complexity: O(n)
// space complexity: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	intervals = append(intervals, newInterval)
	length := len(intervals)
	if length == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i,j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0 ; i< length;i++ {
		if res == nil {
			res = append(res,intervals[i])
			continue
		}

		curLen := len(res)
		last := res[curLen-1]
		if last[1] < intervals[i][0] {
			res = append(res,intervals[i])
		} else {
			last[1] = max(last[1],intervals[i][1])
			res[curLen-1] = last
		}
	}

	return res

}

func max(i,j int) int {
	if i > j {
		return i
	}

	return j
}
func min(i,j int) int {
	if i > j {
		return j
	}

	return i
}

// another concise and straight-forward solution
func insertIntervals(intervals [][]int, newInterval []int) [][]int {
	var i int
	var res [][]int
	start,end,length := newInterval[0],newInterval[1],len(intervals)

	for i < length && intervals[i][1] < start {
		res = append(res,intervals[i])
		i++
	}

	for i < length && intervals[i][0] <= end { // here we mean all the interval that start before newInterval's end
		start = min(intervals[i][0],start)
		end = max(intervals[i][1],end)
		i++
	}

	res = append(res,[]int{start,end})
	for i < length {
		res = append(res,intervals[i])
		i++
	}

	return res
}
