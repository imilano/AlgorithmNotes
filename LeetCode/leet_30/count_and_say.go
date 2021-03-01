package leet_30

import "strconv"

/*
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

    countAndSay(1) = "1"
    countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.

To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

For example, the saying and conversion for digit string "3322251":

Given a positive integer n, return the nth term of the count-and-say sequence.
*/

func countAndSay(n int) string {
	if n < 0 {
		return ""
	}

	res := "1"
	n--
	for n != 0 {
		cur := ""

		for i := 0; i < len(res); i++ {
			cnt := 1
			for i+1 < len(res) && res[i] == res[i+1] {
				cnt++
				i++
			}

			cur += strconv.Itoa(cnt) + string(res[i])
		}

		res = cur
		n--
	}

	return res
}
