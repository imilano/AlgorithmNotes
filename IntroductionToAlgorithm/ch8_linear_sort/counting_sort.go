package ch8_linear_sort

//计数排序算法步骤如下：
//	（1）找出待排序的数组中最大和最小的元素
//	（2）统计数组中每个值为i的元素出现的次数，存入数组bucket的第i项
//	（3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
//	（4）反向填充目标数组：将每个元素i放在新数组的第bucket(i)项，每放一个元素就将bucket(i)减去1
//
//时间复杂度： \Theta (n + k)， n为元素个数，k为元素范围
//空间复杂度：\Theta (n+k). (总空间复杂度n+k, 额外空间复杂度k)

func CountingSort(arr []int) []int {
	max,min := findMaxAndMin(arr)
	return countingSort(arr,max,min)
}

func countingSort(arr []int, maxVal int,minVal int) []int {
	bucketLen := maxVal - minVal + 1
	bucket := make([]int, bucketLen)

	for _,v := range arr {
		bucket[v-minVal] += 1
}

	sortedIndex := 0
	for i := 0;  i< bucketLen; i++ {
		for bucket[i] > 0{
			arr[sortedIndex] = i+minVal
			sortedIndex += 1
			bucket[i]--
		}
	}

	return arr
}

func findMaxAndMin(arr []int) (int,int) {
	max ,min:= arr[0],arr[0]
	for _,v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max,min
}