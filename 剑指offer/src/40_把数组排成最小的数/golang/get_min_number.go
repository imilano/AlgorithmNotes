package golang

import (
	"sort"
	"strconv"
)

// 自定义排序规则，将数字转换为字符串。对于字符串x和y如果x+y > y+x，那么就定义x > y。
func PrintMinNumber( numbers []int ) string {
	// write code here
	//n := len(numbers)
	//if n ==  0 {
	//	return ""
	//}
	//var res string
	//strs := make([]string,n)
	//for i := range  numbers {
	//	strs[i] = strconv.Itoa(numbers[i])
	//}
	//sort.Strings(strs)
	//for i := range strs {
	//	res += strs[i]
	//}
	var res string
	n := len(numbers)
	if n  == 0 {
		return res
	}

	str :=  make([]string,n)
	for i:= range numbers {
		str[i] = strconv.Itoa(numbers[i])
	}

	sort.Slice(str, func(i, j int) bool {
		return str[i] + str[j] < str[j] + str[i]
	})

	for i := range str {
		res += str[i]
	}
	return res
}