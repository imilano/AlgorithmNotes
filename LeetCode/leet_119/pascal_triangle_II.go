package leet_119

/*
Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.
Notice that the row index starts from 0.
Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.

Notice that the row index starts from 0.
*/

func getRow1(rowIndex int) []int {
	res := make([]int,rowIndex+1)
	res[0] = 1
	for i := 1; i< rowIndex+1;i++ {
		for j := i; j >=1; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}


func getRow(rowIndex int) []int {
	var res []int
	res = append(res,1)
	if rowIndex == 0 {
	    return res
	}
	
	res = append(res,1)
	if rowIndex == 1 {
	    return res
	}
	for i:= 2;i <=rowIndex;i++ {
	    tmp := make([]int,len(res))
	    copy(tmp,res)
	    for j:= 1;j< len(res);j++ {
		res[j] = tmp[j-1] + tmp[j]
	    }
    
	    res = append(res,1)
	}
	
	
	return res
    }