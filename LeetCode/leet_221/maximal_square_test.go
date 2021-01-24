package leet_221

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalSquare(matrix))
}
