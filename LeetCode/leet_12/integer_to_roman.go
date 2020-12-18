package leet_12

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

Given an integer, convert it to a roman numeral.
 */

/*
根据题意，可以将相应的映射关系列出为 vector<int>
1000		M
900			CM
500			D
400			CD
100			C
90			XC
50			L
40			XL
10			X
9			IX
5			V
4			IV
1			I

由于题目限制了数字的范围，因而只需要简单的使用贪心算法进行查表即可。
 */

// 每次查表，从最大的数字进行查找，如果相符，减去该数字，然后继续查找下一个。
func intToRomanGreedy(num int) string {
	var res string
	str := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i<len(nums);i++ {
		for num >= nums[i] {
			res += str[i]
			num -= nums[i]
		}
	}

	return res
}

// 比较投机的方法
func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[num%1000/100] + X[num%100/10] + I[num%10]
}
