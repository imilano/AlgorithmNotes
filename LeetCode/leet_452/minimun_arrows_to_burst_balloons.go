package leet_452

import "sort"

func findMinArrowShots(points [][]int) int {
	length := len(points)
	if length == 0 {
		return 0
	}

	// 根据每个区间的end进行排序
	sort.Slice(points, func(i, j int) bool {
		return  points[i][1] < points[j][1]
	})


	// 最少有一个不重叠区间
	start,end,count := 0,points[0][1],1

	for i := 0; i< length ; i++ {
		start = points[i][0]
		if start > end { // 注意此处无等号
			count++
			end = points[i][1]
		}
	}

	return count
}