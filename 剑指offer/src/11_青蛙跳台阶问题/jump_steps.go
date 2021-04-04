package _1_青蛙跳台阶问题


func jumpFloor( number int ) int {
	// write code here
	if number <= 2 {
		return number
	}

	N,NMinusOne,NMinusTwo := 0,2,1
	var count int
	for count < number-2 {
		N = NMinusOne + NMinusTwo
		NMinusTwo = NMinusOne
		NMinusOne = N
		count++
	}

	return N
}
