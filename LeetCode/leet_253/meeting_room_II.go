package leet_253

import (
	"sort"
)

/*
	Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
	find the minimum number of conference rooms required.
*/
// 将开始时间和结束时间分到两个数组中，分别将两个数组进行排序，然后进行遍历，如果开始时间小于结束时间，那么肯定需要再安排一间会议室，故而会议室数量自增；
// 直到有开始时间大于结束时间，那么我们结束时间的位置后移，这意味着前一个结束时间的所空出来的会议室空闲了，会议室可以复用，所以我们只需要简单的前移就好。
func minMeetingRooms(nums [][]int) int {
	var start, end []int
	length := len(nums)
	for i := 0; i < length; i++ {
		start = append(start, nums[i][0])
		end = append(end, nums[i][1])
	}

	sort.Ints(start)
	sort.Ints(end)
	var res, endPos int
	for i := 0; i < length; i++ {
		// 这些会议都需要在结束时间之前开始，意味着我们必须增加会议室数量，否则无法开始会议
		if start != nil && end != nil && start[i] < end[endPos] {
			res++
		} else {
			// 当前会议的开始时间大于会议的结束时间，意味着当前res个会议已经可以结束了，会议结束必然有空出来的会议室，所以我们只需要将结束会议结束向后移动即可
			endPos++
		}
	}

	return res
}

// TODO figure it out
// 每个区间的起始点代表一个区间的开始，会将重叠区域+1，每个区间的结束点代表一个区间的结束，会将重叠区域-1
// 遇到起始时间，映射是正数，则房间数会增加，如果一个时间是一个会议的结束时间，也是另一个会议的开始时间，则映射值先减后加仍为0，
// 并不用分配新的房间，而结束时间的映射值为负数更不会增加房间数
// 将first看做会议的开始时间，m[first\表示first这个时间开始需要的房间数，m[first]++表示这个时间段所需的房间数加1；将second看做会议的结束时间，
// m[second]--表示在second这个时间段所需的房间数减1；如果一个时间点既是一个会议的开始时间，也是一个会议的结束时间，那么相加相减之后恰好为0，无需再添加房间
func minMeetingRooms2(nums [][]int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i][0]]++
		m[nums[i][1]]--
	}

	var rooms, res int
	for _, v := range m {
		rooms += v
		res = max(res, rooms)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
