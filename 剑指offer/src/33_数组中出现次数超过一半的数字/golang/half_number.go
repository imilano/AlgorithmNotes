package golang

func MoreThanHalfNum_Solution( numbers []int ) int {
	// write code here
	n := len(numbers)
	if n < 1 {
		return 0
	}

	var k,v int
	for i := 0; i < n;i++ {
		if v <= 0 {
			k = numbers[i]
			v = 1
		}else if numbers[i] == k {
			v++
		} else if numbers[i] != k {
			v--
		}
	}


	// 进行验证
	var count int
	for i := range numbers{
		if numbers[i] == k {
			count++
		}
	}

	if count <= n / 2 {
		return 0
	}
	return k
}
