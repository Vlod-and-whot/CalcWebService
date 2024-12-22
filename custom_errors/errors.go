package custom_errors

import "errors"

var (
	ErrDivisionByZero    = errors.New("division by zero")
	ErrInvalidExpression = errors.New("invalid expression")
)
