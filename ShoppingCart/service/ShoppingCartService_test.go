package ShoppingCartService

import (
	ShoppingCartModel "REST_API/ShoppingCart/model"
	"reflect"
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

type discountCalculatorTest struct {
	dto      ShoppingCartModel.DiscountCalculatorDTO
	expected map[int]int
}

var discountCalculatorTests = []discountCalculatorTest{
	{ShoppingCartModel.DiscountCalculatorDTO{
		Discount:              map[int]int{1: 0, 8: 0, 18: 0},
		MonthlySpending:       0,
		OrderCountForDiscount: 0,
		GivenAmount:           0,
	}, map[int]int{1: 0, 8: 0, 18: 0}},
	{ShoppingCartModel.DiscountCalculatorDTO{
		Discount:              map[int]int{1: 0, 8: 0, 18: 0},
		MonthlySpending:       10000,
		OrderCountForDiscount: 3,
		GivenAmount:           1000,
	}, map[int]int{1: 10, 8: 10, 18: 15}},
	{ShoppingCartModel.DiscountCalculatorDTO{
		Discount:              map[int]int{1: 0, 8: 0, 18: 0},
		MonthlySpending:       100,
		OrderCountForDiscount: 3,
		GivenAmount:           1000,
	}, map[int]int{1: 0, 8: 10, 18: 15}},
}

func TestDiscountCalculator(t *testing.T) {
	for _, test := range discountCalculatorTests {
		if output := DiscountCalculator(test.dto); reflect.DeepEqual(output, test.expected) {
		} else {
			t.Errorf("Output not equal to expected")
		}
	}
}

type priceCalculationForProductTest struct {
	dto       ShoppingCartModel.PriceCalculationDTO
	expected1 []int
	expected2 []int
}

var priceCalculationForProductTests = []priceCalculationForProductTest{
	{ShoppingCartModel.PriceCalculationDTO{
		Quantity: 1,
		Discount: 10,
		Price:    2000,
		Vat:      1,
		Prices:   nil,
		Vats:     nil}, []int{1800}, []int{18},
	}, {ShoppingCartModel.PriceCalculationDTO{
		Quantity: 2,
		Discount: 15,
		Price:    4000,
		Vat:      18,
		Prices:   nil,
		Vats:     nil}, []int{6800}, []int{1224},
	},
}

func TestPriceCalculationForProduct(t *testing.T) {
	for _, test := range priceCalculationForProductTests {
		prices, vats := test.expected1, test.expected2
		resultPrices, resultVats := PriceCalculationForProduct(test.dto)
		if Equal(vats, resultVats) && Equal(prices, resultPrices) {
		} else {
			t.Errorf("Output not equal to expected")
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
