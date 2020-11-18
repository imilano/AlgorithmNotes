package leet_50

/*
	Implement pow(x, n), which calculates x raised to the power n (i.e. xn).
*/

// Time Limits Exceeded
func myPowRec(x float64, n int) float64 {
	var negative bool
	if n < 0 {
		n = -n
		negative = true
	}

	res := power(x,n/2) * power(x,n/2) * power(x,n%2)
	if negative {
		res = 1/res
	}
	return  res
}

func power(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	if n == 0 {
		return 1
	}

	return power(x,n/2) * power(x,n/2) * power(x,n%2)
}


//---------------------------------
// iterative way
func myPow(x float64, n int) float64 {
	var res float64
	var negative bool
	if n < 0 {
		n = -n
		negative  = true
	}

	if n == 0 {
		return 1
	}

	res = 1
	remaining := n
	for remaining > 0{
		i := 1
		tmp := x
		for i <= remaining {
			tmp *= tmp
			i *= 2
			i *= i
		}
		res *= tmp
		remaining -= i
	}


	if negative {
		res = 1/res
	}

	return res
}


// ------------------------
// divide and conquer
func myPowDC(x float64, n int) float64 {
	if n==0 {
		return 1
	}
	t := myPow(x,n/2)
	if n%2 != 0  {
		if n < 0 {
			return 1/x*t*t
		}else {
			return x*t*t
		}
	} else {
		return t*t
	}
}

//----------------------------------------------
// concise iterative
func myPowConcise(x float64, n int) float64 {
	var ans float64
	if n == 0 {
		return 1
	}

	if n <0 {
		n = -n
		x = 1/x
	}

	ans = 1
	for n != 0 {
		if n & 1 == 1 {
			ans  *= x
		}

		x *= x
		n >>= 1
	}

	return ans
}
