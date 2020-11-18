package leet_11

/*
	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
	n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines,
	which, together with the x-axis forms a container, such that the container contains the most water.
	Notice that you may not slant the container.
*/
func max(a,b int) int {
	if a > b{
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

// time complexity: O(n^2)
// space complexity: O(1)
func brutalForce(height []int) int {
	var maxArea int
	for i:= 0; i< len(height); i++ {
		for j:= i+1; j<len(height);j++ {
			maxArea = max(maxArea, min(height[i],height[j]) * (j-i))
		}
	}

	return maxArea
}

// 使用两个指针left和right，初始时，一个指针指向最左边，另一个指针指向最右边；计算此时二者的maxArea；然后，如果左边的高度小于右边的高度，那么就向后移动left；否则向前移动right
// 为什么只是移动高度小的那个指针，而不是移动高度大的那个指针呢？很显然，在两个高度中，较小者决定了面积的大小，如果我们移动较大的指针，那么不仅宽度在减少，高度也在减小，如此必然不能得到更大的面积；
// 而如果我们移动高度较小的指针，那么虽然我们的宽度减小了，但是可以通过高度得到补偿，所以，移动较矮的指针显然更为有效。
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func maxAreaWithTwoPointer(height []int) int {
	maxArea,right,left := 0,len(height)-1,0

	for left < right {
		maxArea = max(maxArea,min(height[left],height[right]) * (right-left))
		if height[left] < height[right] {
			left++
		} else {
			right --
		}
	}

	return maxArea
}

func maxArea(height []int) int {
	// brutal force
	return brutalForce(height)
}