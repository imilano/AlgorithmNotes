package leet_271

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Printf("%s\n", "\\0")

	str := "aa\\0bb\\0"
	fmt.Println(len(str))
	n := len(str)
	for i := 0; i < len(str); {
		if i+1 < n && str[i] == '\\' && str[i+1] == '0' {
			i += 2
			continue
		}
		fmt.Println(strconv.Itoa(int(str[i])))
		i++

	}
}

func TestEncodeAndDecode(t *testing.T) {
	strs := [][]string{
		[]string{"aa", "bb", "cc", "00", "a878d"},
		[]string{"002", ",.+-\fdfd", "cddgadgf.."},
	}

	for _, str := range strs {
		e := encode(str)
		d := decode(e)

		if !reflect.DeepEqual(d, str) {
			t.Errorf("test fail: want %v, get %v", str, d)
		}
	}
}
