package Custom_Errors

import "errors"

var (
	ErrDivisionByZero    = errors.New("division by zero")
	ErrInvalidExpression = errors.New("invalid expression")
)
