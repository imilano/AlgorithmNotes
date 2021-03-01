package leet_17

/*
	Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
	Return the answer in any order.
	A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/

//---------------------------------------------
// recursive solution
func helper(digits string, index int, combination string, result *[]string, dict map[uint8][]string) {
	length := len(digits)
	if index > length-1 {
		*result = append(*result, combination)
		return
	}

	c := digits[index]
	for i := range dict[c] {
		helper(digits, index+1, combination+dict[c][i], result, dict)
	}
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	// numbers could be replaced by slice
	numbers := map[uint8][]string{
		'1': {""},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	helper(digits, 0, "", &res, numbers)
	return res
}

//--------------------------------
// iterative solution
func letterCombinations2(digits string) []string {
	//res := make([]string,4)
	res := make([]string, 1)
	length := len(digits)
	if length == 0 {
		return nil
	}

	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for i := 0; i < length; i++ {
		var cur []string
		str := dict[int(digits[i]-'0')]
		for j := 0; j < len(str); j++ {
			for _, s := range res { // res 不能为nil，否则无法进入这个循环，故而res初始化时size应该为1
				cur = append(cur, s+string(str[j]))
			}
		}

		res = cur
	}

	return res
}
