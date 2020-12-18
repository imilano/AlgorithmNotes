package ch6_heap_sort

// heap sort , using max heap for demonstration
// index start from 1,not 0

// adjust heap
// time complexity: O(lgn)
func MaxHeapify(nums []int, index int) {
	// its left child and right child
	left, right := 2*index, 2*index+1
	largest := index

	// find the largest number between its left and right child
	if left <= len(nums) && nums[left] > nums[index] {
		largest = left
	}

	if right <= len(nums) && nums[right] > nums[largest] {
		largest = right
	}

	// if index is largest, then we did not need to swap ,then heap is already sorted.
	// Otherwise we swap  index and largest, then continue to adjust
	if largest != index {
		// swap element
		nums[index], nums[largest] = nums[largest], nums[index]
		MaxHeapify(nums, largest)
	}
}

// build heap
func BuildMaxHeap(nums []int) {
	length := len(nums)

	for i := length / 2; i > 0; i-- {
		MaxHeapify(nums, i)
	}
}

// heap sort
// time complexity: O(nlgn)
func HeapSort(nums []int) {
	// first build a heap
	BuildMaxHeap(nums)

	// then swap the first element with the last, the adjust the heap, then reduce the length by 1
	// repeat until has only one element left
	length := len(nums)
	for i := len(nums); i > 1; i-- {
		nums[length], nums[1] = nums[1], nums[length]
		length--
		MaxHeapify(nums, 1)
	}
}
