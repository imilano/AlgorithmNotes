package leet_248

/*
A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Write a function to count the total strobogrammatic numbers that exist in the range of low <= num <= high.

Example:

Input: low = "50", high = "100"
Output: 3
Explanation: 69, 88, and 96 are three strobogrammatic numbers.
Note:
Because the range might be a large number, the lowand high numbers are represented as string.
*/

//还是继续之前采用过的思路，但是需要排除一些情况：
//- 所有长度不介于low和high之间的情况
//- 长度介于low和high之间，但是以0开头
//- 长度介于low和high之间，但是比low小或者比high大
func strobogrammaticInRange(low, high string) int {
	var res int
	find(low, high, "", &res)
	find(low, high, "0", &res)
	find(low, high, "1", &res)
	find(low, high, "8", &res)
	return res
}

func find(low, high, path string, res *int) {
	if len(path) >= len(low) && len(path) <= len(high) {
		if len(path) == len(high) && path > high { // 排除那些比high大的元素
			return
		}

		// 排除比low小并且以0开头的元素
		if (len(path) > 1 && path[0] == '0') != true && (len(path) == len(low) && path < low) != true {
			*res++
		}
	}

	if len(path)+2 > len(high) {
		return
	}

	find(low, high, "0"+path+"0", res)
	find(low, high, "1"+path+"1", res)
	find(low, high, "8"+path+"8", res)
	find(low, high, "6"+path+"9", res)
	find(low, high, "9"+path+"9", res)
}
