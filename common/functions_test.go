package common

import (
	"reflect"
	"testing"
)

type discountCalculatorTest struct {
	discount              map[int]int
	monthlySpending       int
	vATTypes              int
	orderCountForDiscount int
	givenAmount           int
	expectedDiscount      map[int]int
}
type sumTest struct {
	slice    []int
	expected int
}

var sumTests = []sumTest{
	{[]int{1, 2, 3, 4}, 10},
	{[]int{9, 3, 4}, 16},
	{[]int{3920, 80}, 4000},
	{[]int{10, 20, 40}, 70},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		if output := SumOfIntSlice(test.slice); output != test.expected {
			t.Errorf("Output not equal to expected")
		}
	}
}

var discountCalculatorTests = []discountCalculatorTest{
	{map[int]int{1: 1, 8: 1, 18: 1}, 0, 3, 0, 0, map[int]int{1: 1, 8: 1, 18: 1}},
	{map[int]int{1: 1, 8: 1, 18: 1}, 10000, 3, 7, 5000, map[int]int{1: 10, 8: 10, 18: 15}},
}

func TestDiscountCalculator(t *testing.T) {
	for _, test := range discountCalculatorTests {
		output2 := DiscountCalculator(test.discount, test.monthlySpending, test.vATTypes, test.orderCountForDiscount, test.givenAmount)
		if reflect.DeepEqual(output2, test.expectedDiscount) {
		} else {
			t.Errorf("Output not equal to expected")
		}
	}
}
