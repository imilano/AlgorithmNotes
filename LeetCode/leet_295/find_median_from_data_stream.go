package leet_295

import "sort"

/*
Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.


Example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2


Follow up:

If all integer numbers from the stream are between 0 and 100, how would you optimize it?
If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?
*/

// 方法1，采用排序的方法，由于每次findMedian都要对数组进行排序，所以效率较低。
// 方法2，采用最大堆和最小堆的方法。将数据分为两部分，左边是较小的部分，存入最大堆，右边是大的部分，存入最小堆。当插入元素的时候，优先插入最小堆，
//此时如果最小堆的元素总量比最大堆大2，则将最小堆堆顶元素插入最大堆，然后删除最小堆对堆顶元素。在查找时，如果两个堆的元素个数相等，则分别取出堆顶元素，
//取中值返回即可；否则返回最小堆的堆顶元素。
type MedianFinder struct {
	nums []int
	n int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		nums: []int{},
		n: 0,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	this.nums = append(this.nums,num)
	this.n++
}


func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.nums)

	if this.n % 2 == 1 {
		return float64(this.nums[this.n/2])
	} else {
		return float64(this.nums[this.n/2]+this.nums[this.n/2-1])/2
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */