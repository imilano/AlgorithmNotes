package leet_295

import (
	"fmt"
	"testing"
)

func TestFindMedian(t *testing.T) {
	m := Constructor2()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
}
