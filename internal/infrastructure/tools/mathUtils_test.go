package tools

import (
	"fmt"
	"testing"
)

func TestCompareNumberStrings(t *testing.T) {

	a := "0"
	b := "4000"
	value, err := CompareNumberStrings(a, b)
	if err != nil {

		t.Error(err.Error())
	}

	fmt.Println(value)
}

func TestSubtractStringNumbers(t *testing.T) {

	a := "47.99999999999999"
	b := "5.2"
	n := 1
	value, err := SubtractStringNumbers(a, b, float64(n))
	if err != nil {

		t.Error(err.Error())
	}

	fmt.Println(value)
}
