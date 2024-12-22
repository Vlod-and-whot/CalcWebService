package CalculationTest

import (
	"CalculationWebService/Packages/Calculation"
	"CalculationWebService/Packages/Custom_Errors"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  error
	}{
		{"(1 + 2) / 0", 0, Custom_Errors.ErrDivisionByZero},
		{"+ 5", 0, Custom_Errors.ErrInvalidExpression},
		{"12 - fdsfdf", 0, Custom_Errors.ErrInvalidExpression},
		{"1 + 2", 3, Custom_Errors.ErrInvalidExpression},
		{"4 - 1 *", 0, Custom_Errors.ErrInvalidExpression},
		{"7 * 8", 56, nil},
		{"123 - 100", 23, nil},
		{"(33 / 11) * 3", 9, nil},
		{"2 + 6 * 2", 14, nil},
		{"3 * 7 - 5", 16, nil},
		{"54 / (6 + 3)", 6, nil},
		{"2 + 2 * 2", 6, nil},
		{"2 + 2 / 2", 1, nil},
		{"2 / 1", 2, nil},
		{"2 + 3 * (4 + 5)", 29, nil},
		{"2 + 3 - 1", 4, nil},
		{"(3 - 2) * (3 + 2)", 5, nil},
		{"4 + 2 * 2 * 3 + 9", 25, nil},
		{"1 - 1", 0, nil},
	}

	for _, test := range tests {
		result, err := Calculation.Calc(test.expression)
		if result != test.expected || (err != nil && err.Error() != test.expectErr.Error()) {
			t.Errorf("Calc(%q) = %v, %v; want %v, %v", test.expression, result, err, test.expected, test.expectErr)
		}
	}
}
