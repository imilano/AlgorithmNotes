package leet_169

import "sort"

/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array
*/

// slow solution
// time complexity: O(nlogn)
// space complexity: O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// time complexity: O(n)
// space complexity: O(n)
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	length := 0
	for _,v:= range nums {
		m[v]++
		length++
	}

	for v,t := range m {
		if t > length/2 {
			return v
		}
	}

	return 0
}

// Moore voting
// time complexity: O(n)
// space complexity: O(1)

//摩尔投票法，类似于相互抵消的思想，只要对数组中有过半数的数字时才能使用。假设当前数字是那个过半数，如果下一个数字跟当前数字相同，那么将计数加加；如果不同，则将计数减减；
//如果计数减到了0，那么将当前过半数更换为下一个数字。这样，最后出现的那个数字就是我们所需要的过半数。
func majorityElement3(nums []int) int {
	var res,cnt int
	for _,v := range nums {
		if cnt == 0 {
			res = v
			cnt++
		} else {
			if v == res {
				cnt ++
			} else {
				cnt--
			}
		}
	}

	return res
}

// bit manipulation
func majorityElement4(nums []int) int {
	//var ones,zeros,res int
	var res int32
	var ones,zeros int
	length := len(nums)

	 for i := 0; i< 32;i++ {
	 	ones,zeros = 0,0
	 	for _,v := range nums {
	 		if ones > length/2 || zeros > length/2 {
	 			break
			}
	 		if ((v >>i) & 1) == 1 {
	 			ones++
			} else {
				zeros++
			}
		}

		if ones > zeros {
			res |= 1 << i
		}
	 }

	 //return res
	 return int(res)
}