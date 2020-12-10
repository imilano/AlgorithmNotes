package leet_204

import "math"

/*
Count the number of prime numbers less than a non-negative number, n.
Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

// Brutal Force
// Time Limit Exceeded
func isPrime(n int) bool {
	// n/2 does not equals to the root of n
	root := int(math.Sqrt(float64(n)))
	for i := 2; i<=root;i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	var res int
	for i := 2; i< n;i++ {
		if isPrime(i) {
			res++
		}
	}


	return res
}


// 对于每一个素数，标记每一个素数的倍数。由于所求为不大于n，故下标只需要到n-1，故而只需要开辟n个空间。
func countPrimes2(n int) int {
	var res int
	m := make(map[int]int,n)

	for i := 2; i<n;i++ {
		if m[i] != 0 {
			continue
		}

		res++

		// 标记
		for j := 2; j*i < n;j++ {
			m[i*j] = 1
		}
	}

	return res
}
