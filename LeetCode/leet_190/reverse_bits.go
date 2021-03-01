package leet_190

/*
 Reverse bits of a given 32 bits unsigned integer.

Note:

Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.
Follow up:

If this function is called many times, how would you optimize it?
*/

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 1; i <= 32; i++ {
		bit := (num >> (i - 1)) & 0x1
		res += bit << (32 - i)
		//res += ((num >>(i-1)) & 1) << (32-i)
	}

	return res
}
