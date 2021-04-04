package _0_斐波那契数列

// F(1) = 1, F(0) = 0, F(n) = F(n-1) + F(n-2)
func Fibonacci( n int ) int {
	// write code here
	if n < 2 {
		return n
	}

	var count int
	N,NMinusOne,NMinusTwo := 0,1,0
	for count <  n-1 {
		N = NMinusOne + NMinusTwo
		NMinusTwo = NMinusOne
		NMinusOne = N
		count++
	}
	return N
}

func FibRec(n int) int {
	if n < 2 {
		return n
	}

	return FibRec(n-1) + FibRec(n-2)
}

