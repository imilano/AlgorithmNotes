package leet_17

/*
	Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
	Return the answer in any order.
	A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
 */

func helper(digits string,index int,combination string,result *[]string,dict map[uint8][]string) {
	length := len(digits)
	if index > length-1 {
		*result = append(*result,combination)
		return
	}

	c := digits[index]
	for i := range dict[c] {
		helper(digits,index+1,combination+dict[c][i],result,dict)
	}
}


func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	numbers := map[uint8][]string{
		'1':{""},
		'2':{"a","b","c"},
		'3':{"d","e","f"},
		'4':{"g","h","i"},
		'5':{"j","k","l"},
		'6':{"m","n","o"},
		'7':{"p","q","r","s"},
		'8':{"t","u","v"},
		'9':{"w","x","y","z"},

	}
	helper(digits,0,"",&res,numbers)
	return res
}
