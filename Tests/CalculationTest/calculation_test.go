package CalculationTest

import (
	"CalculationWebService/Packages/Calculation"
	"CalculationWebService/Packages/Custom_Errors"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
		expectErr  error
	}{
		{"error", "(1 + 2) / 0", 0, Custom_Errors.ErrDivisionByZero},
		{"error", "+ 5", 0, Custom_Errors.ErrInvalidExpression},
		{"error", "12 - fdsfdf", 0, Custom_Errors.ErrInvalidExpression},
		{"error", "1 + 2", 3, Custom_Errors.ErrInvalidExpression},
		{"error", "4 - 1 *", 0, Custom_Errors.ErrInvalidExpression},

		{"simple", "7 * 8", 56, nil},
		{"simple", "123 - 100", 23, nil},
		{"simple", "(33 / 11) * 3", 9, nil},
		{"simple", "2 + 3 - 1", 4, nil},
		{"simple", "2 / 1", 2, nil},
		{"simple", "1 - 1", 0, nil},

		{"priority", "2 + 6 * 2", 14, nil},
		{"priority", "3 * 7 - 5", 16, nil},
		{"priority", "54 / (6 + 3)", 6, nil},
		{"priority", "2 + 2 * 2", 6, nil},
		{"priority", "2 + 2 / 2", 1, nil},
		{"priority", "2 + 3 * (4 + 5)", 29, nil},
		{"priority", "(3 - 2) * (3 + 2)", 5, nil},
		{"priority", "4 + 2 * 2 * 3 + 9", 25, nil},

		{"hard", "1 * 2 * 3 * 4 / 5 - 6 + 8 - (18 - 9) * 5 / 3 - 12 + 128 * (5 - 3) * (5 + 3)", 2027.8, nil},
		{"hard", "1234567890 * 9876543210 / 1234567890 * 2 / 4", 4938271605.0, nil},
	}

	for _, test := range tests {
		result, err := Calculation.Calc(test.expression)
		if result != test.expected || (err != nil && err.Error() != test.expectErr.Error()) {
			t.Errorf("Calc(%q) = %v, %v; want %v, %v", test.expression, result, err, test.expected, test.expectErr)
		}
	}
}
