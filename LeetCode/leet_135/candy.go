package leet_135

/*
There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?
*/

//两趟扫描。第一次从左到右，当前rating大于前一个的rating时，当前的rating分得的蛋糕要比前一个多1，确保从左往右高等级的孩子分到的糖果更多。
//然后从右往左扫描，如果当前rating大于后一个的rating，并且当前孩子的糖果小于后一个孩子的糖果时，当前孩子所分糖果在后一个的基础上加1。
func candy(ratings []int) int {
	length := len(ratings)
	candy := make([]int, length)
	for i := range candy {
		candy[i] = 1 // at least 1 candy for every child
	}

	// 注意i的写法
	for i := 0; i < length-1; i++ {
		if ratings[i+1] > ratings[i] {
			candy[i+1] = candy[i] + 1
		}
	}

	// 注意i的写法
	for i := length - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candy[i-1] = max(candy[i-1], candy[i]+1)
		}
	}

	var sum int
	for _, v := range candy {
		sum += v
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
