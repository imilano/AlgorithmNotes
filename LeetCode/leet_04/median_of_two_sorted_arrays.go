package leet_04

import "math"

/*
 *	Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
 *  Follow up: The overall run time complexity should be O(log (m+n)).
 */

/*
	题解见同文件夹markdown。
 */

func getMax(x,y int) int {
	if x > y {
		return x
	}

	return y
}

func getMin(x,y int) int {
	if x > y {
		return y
	}

	return x
}

// time complexity: O(log(min(m,n))
// detail see markdown
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2,nums1)
	}

	x,y := len(nums1),len(nums2)
	low,high := 0,x
	var maxLeftX,maxLeftY,minRightX,minRightY int
	for low <= high {
		partitionX := (low+high)/2
		partitionY := (x+y+1)/2 - partitionX

		// edge condition
		if partitionX == 0 {
			maxLeftX = math.MinInt32
		} else {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX == x {
			minRightX = math.MaxInt32
		} else {
			minRightX = nums1[partitionX]
		}

		if partitionY == 0 {
			maxLeftY = math.MinInt32
		} else {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY == y {
			minRightY = math.MaxInt32
		} else {
			minRightY = nums2[partitionY]
		}

		// We found the correct place
		// If the total length is odd number, then we choose the  bigger one between maxLeftX and maxLeftY
		// If the total length is even number, then we choose the the average between the max of maxLeft* and minRight*
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0{
				return (float64(getMax(maxLeftY,maxLeftX)) + float64(getMin(minRightY,minRightX)))/2
			} else {
				return float64(getMax(maxLeftX,maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// We go far right, we need move left
			high = partitionX -1
		} else  {
			// We go far left, we need move right
			low = partitionX + 1
		}
	}
	return 0.0
}


