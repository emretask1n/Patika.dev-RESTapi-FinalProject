package ShoppingCartService

import (
	"testing"
)

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
