package leet_43

import "strconv"

/*
	Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2,
	also represented as a string.
	Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

// Original solution: brutal force
// addMul return the sum of the integer
func addMul(a,b string) string {
	// suppose a is of shortest length
	if len(a) > len(b) {
		return addMul(b,a)
	}

	// we always suppose a is smaller one, so ,if a is empty, we just need to return b
	if len(a) == 0 {
		return b
	}

	la,lb := len(a)-1,len(b)-1
	var carry uint8
	var res string
	for la >= 0{
		sum := a[la] - '0' + b[lb]-'0' + carry
		carry = sum/ 10
		sum %= 10
		res = strconv.Itoa(int(sum)) + res

		la--
		lb--
	}

	for lb >= 0 {
		sum := b[lb] - '0' + carry
		carry = sum/ 10
		sum %= 10
		res = strconv.Itoa(int(sum)) + res

		lb--
	}

	if carry > 0 {
		res = strconv.Itoa(int(carry)) + res
	}

	return res
}

// getMul return the product of two integer, we suppose a is bigger than b
func getMul(a string, b uint8) string{
	if b == '0' {
		return "0"
	}
	// suppose b is smallest one
	la := len(a)-1

	var res string
	var carry uint8
	for la >= 0 {
		pro := (a[la] - '0')*(b-'0') + carry
		carry = pro/ 10
		pro %= 10

		res = strconv.Itoa(int(pro)) + res
		la--
	}

	if carry > 0 {
		res = strconv.Itoa(int(carry)) + res
	}


	return res
}


func multiply(num1 string, num2 string) string {
	// we always suppose nums1 is the smaller one
	if len(num1) >  len(num2) {
		return multiply(num2,num1)
	}

	var res string
	la := len(num1) -1
	for la  >= 0 {
		cur := getMul(num2,num1[la])
		res = addMul(cur,res)

		la--
		num2 += "0"
	}

	return res
}

//
//func multiplyQui(num1 string, num2 string) string {
//	//if len(num1) > len(num2) {
//	//	return multiply(num2,num1)
//	//}
//
//	//if len(num1) == 1 && num1[0] == '0' {
//	//	return "0"
//	//}
//
//	la,lb := len(num1),len(num2)
//	res := make([]uint8,la+lb)
//	for i := lb-1;i>=0;i-- {
//		for j := la-1; j >= 0; j-- {
//			cur := (num2[i]-'0')*(num1[j]-'0')
//			p1 := i+j
//			p2 := i+j +1
//			sum := cur + res[p2]
//
//			res[p2] = sum % 10
//			res[p1] += sum/10
//		}
//	}
//
//	index := la+lb-1
//	for i := range res {
//		if res[i] != 0 {
//			index = i
//			break
//		}
//	}
//	res = res[index:]
//
//	var s string
//	for _,v := range res {
//		s += strconv.Itoa(int(v))
//	}
//
//	return s
//}

// `num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`
func multiplyQuick(num1 string, num2 string) string {
	la,lb := len(num1),len(num2)
	res := make([]uint8,la+lb)
	for i := lb-1;i>=0;i-- {
		for j := la-1; j >= 0; j-- {
			cur := (num2[i]-'0')*(num1[j]-'0')
			p1 := i+j
			p2 := i+j +1
			sum := cur + res[p2]

			res[p2] = sum % 10
			res[p1] += sum/10
		}
	}

	// skip head 0, set index to la+lb-1 to avoid all the bit is zero, in this case, we just need to return the last zero
	index := la+lb-1
	for i := range res {
		if res[i] != 0 {
			index = i
			break
		}
	}
	res = res[index:]

	var s string
	for _,v := range res {
		s += strconv.Itoa(int(v))
	}

	return s
}
