package leet_296

import "sort"

/*
A group of two or more people wants to meet and minimize the total travel distance. You are given a 2D grid of values 0 or 1, where each 1 marks the home of someone in the group. The distance is calculated using Manhattan Distance, where distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|.

Example:

Input:

1 - 0 - 0 - 0 - 1
|   |   |   |   |
0 - 0 - 0 - 0 - 0
|   |   |   |   |
0 - 0 - 1 - 0 - 0

Output: 6

Explanation: Given three people living at (0,0), (0,4), and (2,2):
             The point (0,2) is an ideal meeting point, as the total travel distance
             of 2+2+2=6 is minimal. So return 6.
Hint:

Try to solve it in one dimension first. How can this solution apply to the two dimension case?
*/


// 按照题目提示，先从1维的情况来分析，比如[A.......B]，很明显，p只能在A和B中间内即可，如果不在A与B之间，那么距离一定大于线段AB；同理对于
//[A.....B.....C......D]，P必须位于B与C之间，这样总的最小距离就是AD和BC之和，一旦P不位于BC之间，那么其距离一定大于AD+BC。
// 总结就是，给坐标排序，然后用最大坐标减去最小坐标，然后依次向内收缩，继续计算；对于二维数组，单独计算维度，然后将结果相加即可。
func minTotalDistance(grid [][]int) int{
	var res int
	var rows,cols []int
	for i := 0; i < len(grid);i++ {
		for j := 0; j < len(grid[0]);j++ {
			if grid[i][j] == 1{
				rows = append(rows,i)
				cols = append(cols,j)
			}
		}
	}

	i,j := 0,len(rows)-1
	for i < j {
		res += rows[j]-rows[i]
		j--
		i++
	}

	sort.Ints(cols)
	i,j  = 0,len(cols)-1
	for i < j {
		res += cols[j]-cols[i]
		i++
		j--
	}

	return  res
}
