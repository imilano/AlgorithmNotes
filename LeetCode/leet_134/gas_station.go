package leet_134

/*

There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

Note:

If there exists a solution, it is guaranteed to be unique.
Both input arrays are non-empty and have the same length.
Each element in the input arrays is a non-negative integer.
*/

//dp. 从start开始遍历，计算到达i点时刻的余量之和debt，如果余量一直大于0，那么可以一直向后遍历；如果余量和小于0，那么说明从起点开始不能走到当前点，那么直接从i+1开始进行遍历。
//为什么从i+1开始遍历呢？如果从start走到i点时，debt刚好小于0的话，从start到start+1，debt肯定是大于0的，debt加上一个大于0的数之后还比0小，说明从start+1肯定到不了i，同理对其他点可得。
//故而应该从i+1开始进行遍历。另外，很明显，如果从开始遍历到结束，计算的总数小于0的话，整个环的净余量都是小于0的，肯定不能走完。
func canCompleteCircuit(gas []int, cost []int) int {
	var total,debt int
	var start int
	for i := 0; i<len(gas);i++ {
		debt += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if debt < 0 {
			start = i + 1
			debt = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}