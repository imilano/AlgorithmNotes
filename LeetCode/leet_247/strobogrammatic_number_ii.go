package leet_247

/*
A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Find all strobogrammatic numbers that are of length = n.

Example:

Input:  n = 2
Output: ["11","69","88","96"]
*/

func findStrobogrammatic(n int) []string {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		return []string{"0", "1", "8"}
	}

	var res []string
	for _, v := range findStrobogrammatic(n - 2) {
		res = append(res, "1"+v+"1", "8"+v+"8", "6"+v+"9", "9"+v+"6", "0"+v+"0")
	}

	return res
}
