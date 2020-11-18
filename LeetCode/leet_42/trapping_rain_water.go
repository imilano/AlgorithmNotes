package leet_42

/*
	Given n non-negative integers representing an elevation map where the width of each bar is 1,
	compute how much water it can trap after raining.
 */

// tool function
func count(height []int, left, right int) int {
	width := right - left -1
	area := height[left] * width
	for i := left +1; i< right ;i ++ {
		area -= height[i]
	}

	return area
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a,b int) int {
	if a > b {
		return b
	}

	return a
}

// ---------------------------------------------------------------------
// dynamic programming use two array to reduce time complexity, this is a quite clever solution
// time complexity:O(n)
// space complexity: O(n)
func trapWith2Scan(height []int) int {
	var area int
	length := len(height)
	if length <= 1{
		return area
	}

	leftMax,rightMax := make([]int, length),make([]int,length)
	// get max from left to right
	leftMax[0] = height[0]
	for i := 1; i< length;i++ {
		leftMax[i] = max(leftMax[i-1],height[i])
	}

	// get max from right to left
	rightMax[length-1] = height[length-1]
	for i := length-2; i >= 0;i-- {
		rightMax[i] = max(height[i],rightMax[i+1])
	}

	// take care, here we use min, not max
	for i := 0; i< length;i++ {
		area += min(leftMax[i],rightMax[i]) - height[i]
	}


	return area
}

//---------------------------------------------
// use two pointer
// time complexity: O(N)
// space complexity: O(1)
func trapWith2Pointer(height []int) int {
	var res int
	if len(height) <= 1 {
		return res
	}

	left,right := 0,len(height)-1
	leftMax,rightMax := height[0],height[len(height)-1]

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else  {
				res += rightMax - height[right]
			}

			right--
		}
	}


	return res
}


// ------------------------------------------------
// Original wrong solution
func trap(height []int) int {
	var area int
	length := len(height)
	if length <= 1{
		return area
	}


	maxHeight := 0
	for i := range height {
		if height[i] > height[maxHeight] {
			maxHeight = i
		}
	}

	lmin,lmax := maxHeight,maxHeight
	for lmin >= 0 {
		if lmin-1 >= 0 && !(height[lmin] >= height[lmin-1] && height[lmin] >= height[lmin+1]) {
			lmin--
			continue
		}
		area += count(height,lmin,lmax)
		lmax =  lmin
		lmin--
	}

	rmin,rmax := maxHeight,maxHeight
	for rmax < length {
		if rmax+1 < length && height[rmax] >= height[rmax+1] && height[rmax] >= height[rmax-1] {
			rmax++
			continue
		}

		area += count(height,rmin,rmax)
		rmin = rmax
		rmax++
	}

	return area
}