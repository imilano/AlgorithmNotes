package golang

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrintMinNumber(t *testing.T) {
	strs := [][]string{
		{"3","32","321"},
		{"11","22","34"},
	}

	for i := range strs{
		sort.Strings(strs[i])
		fmt.Println(strs[i])
		var res string
		for j := range  strs[i] {
			res += strs[i][j]
		}

		fmt.Println(res)
	}
}
