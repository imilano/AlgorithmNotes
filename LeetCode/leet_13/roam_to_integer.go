package leet_13

/*
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.
*/

/*
注意这几个规则：
1、相同的数字连写，所表示的数等于这些数字相加得到的数，如：Ⅲ = 3；
2、小的数字在大的数字的右边，所表示的数等于这些数字相加得到的数， 如：Ⅷ = 8；Ⅻ = 12；
3、小的数字，（限于Ⅰ、X 和C）在大的数字的左边，所表示的数等于大数减小数得到的数，如：Ⅳ= 4；Ⅸ= 9；
*/

// 使用map，将罗马数字字母转换为对应的整数值
func romanToInt2(s string) int {
	var res int
	roman := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)
	var val int
	for i := 0; i < length; i++ {
		val = roman[s[i]]
		if i == length-1 || roman[s[i]] >= roman[s[i+1]] {
			res += val
		} else {
			res -= val
		}
	}

	return res
}

func romanToInt(s string) int {
	var res int
	roman := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)
	for i := 0; i < length; {
		if s[i] == 'I' && i+1 < length && s[i+1] == 'V' {
			res += 4
			i += 2
			continue
		}
		if s[i] == 'I' && i+1 < length && s[i+1] == 'X' {
			res += 9
			i += 2
			continue
		}

		if s[i] == 'X' && i+1 < length && s[i+1] == 'L' {
			res += 40
			i += 2
			continue
		}

		if s[i] == 'X' && i+1 < length && s[i+1] == 'C' {
			res += 90
			i += 2
			continue
		}

		if s[i] == 'C' && i+1 < length && s[i+1] == 'D' {
			res += 400
			i += 2
			continue
		}

		if s[i] == 'C' && i+1 < length && s[i+1] == 'M' {
			res += 900
			i += 2
			continue
		}
		res += roman[s[i]]
		i++
	}

	return res
}
