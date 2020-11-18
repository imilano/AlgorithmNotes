package leet_56

import "sort"

/*
	Given a collection of intervals, merge all overlapping intervals.

	Example 1:
	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
*/

//type gap [][]int
//
//func (s gap) Less(i,j int) bool {
//	return s[i][0] < s[j][0]
//}
//
//func (s gap)Swap(i,j int) {
//	s[i],s[j] = s[j],s[i]
//}
//
//func (s gap) Len() int {
//	return len(s)
//}
//

// time complexity: O(nlogn)
// space complexity: O(1)
// First we sort the intervals by their start value in ascending order
// Then we insert  interval to res, we do  described as follows:
// IF the current interval start after the previous interval ends, then we just insert it to res;
// else if current interval start before the previous interval ends, then we update the previous interval ends to the max value of end.
func max(i,j int)  int {
	if i > j {
		return i
	}

	return j
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]  < intervals[j][0]
	})

	for _,interval := range intervals {
		if res == nil {
			res = append(res, interval)
			continue
		}

		last := res[len(res)-1]
		if last[1] < interval[0] {
			res = append(res, interval)
		} else {
			last[1] = max(last[1], interval[1])
			res[len(res)-1] = last
		}
	}

	return res
}