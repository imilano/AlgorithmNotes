package ch7_quick_sort

import (
	"math/rand"
	"time"
)

// 快排的性能取决于划分的好坏，如果划分的好，那么快排的性能就能达到O(nlgn)，如果划分的不好，就可能出现最坏情况O(n^2)
func QuickSort(nums []int, left,right int) {
	if left < right {
		div := Partition(nums,left,right)
		QuickSort(nums,left,div-1)
		QuickSort(nums,div+1,right)
	}
}

func Partition(nums []int, left,right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i<= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums,i,index)
			index++  // now index points to a value which larger than pivot
		}
	}

	// notice  that we swap to index -1, not index
	swap(nums,pivot,index-1)
	return index-1
}

func swap(nums []int, a,b int) {
	nums[a],nums[b] = nums[b],nums[a]
}

// Improve QuickSort to reduce time complexity
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func RandomPartition(nums []int, left,right int) int {
	r := random(left,right)
	swap(nums,left,r)
	return Partition(nums,left,right)
}


func main() {
	nums := []int {1,2,4,4,5,2,5,2,99}
	QuickSort(nums,0,len(nums)-1)
}
