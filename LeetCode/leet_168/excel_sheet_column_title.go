package leet_168

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
*/

//-------------------------------------
// 查字典的方式，问题在于对正好为26整数倍值的那些n，无法得到Z。 解决办法是使用小trick：比如对于26来说，先减一得到25，再对26取余得到25，再对
// 取余得到的数加1得到26.这种办法对于其他不能整数26的数也成立
func convertToTitle(n int) string {
	m := map[int]string {
		1:"A", 2:"B", 3:"C", 4:"D",5:"E",6:"F",7:"G",8:"H",9:"I",10:"J",
		11:"K",12:"L",13:"M",14:"N",15:"O",16:"P",17:"Q",18:"R",19:"S",20:"T",
		21:"U",22:"V",23:"W",24:"X",25:"Y",26:"Z",
	}

	var res string
	for n != 0 {
		// use trick
		n--
		remainder := n %26 + 1
		quotient := n / 26
		res = m[remainder] + res
		if quotient > 26 {
			n = quotient
		} else {
			res = m[quotient] + res
			return res
		}
	}

	return res
}

//---------------------------
// recursive
func convertToTitleOneLine(n int) string {
	if n  == 0 {
		return ""
	}

	return convertToTitle((n-1)/26) + string((n-1)%26 + 'A')
}
