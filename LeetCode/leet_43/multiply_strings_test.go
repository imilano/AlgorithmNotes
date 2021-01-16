package leet_43

import (
	"fmt"
	"testing"
)

func TestAddNum(t *testing.T) {
	test := []struct {
		a   string
		b   string
		sum string
	}{
		{"111", "222", "333"},
		{"101", "99", "200"},
		{"900", "100", "1000"},
	}

	for _, s := range test {
		sum := addMul(s.a, s.b)
		if sum != s.sum {
			fmt.Printf("%v %s\n", s, sum)
		}
	}
}

func TestGetMul(t *testing.T) {
	test := []struct {
		a string
		b uint8
		c string
	}{
		{"111", '1', "111"},
		{"200", '5', "1000"},
		{"300", '0', "0"},
	}

	for _, s := range test {
		pro := getMul(s.a, s.b)
		if pro != s.c {
			fmt.Printf("%v %s\n", s, pro)
		}
	}
}

func TestMultiply(t *testing.T) {
	test := []struct {
		a   string
		b   string
		sum string
	}{
		{"123", "456", "56088"},
	}

	for _, s := range test {
		pro := multiply(s.b, s.a)
		if pro != s.sum {
			fmt.Printf("%v %s\n", s, pro)
		}
	}
}

func TestMultiplyQui(t *testing.T) {
	multiplyQui("7", "0")
}
