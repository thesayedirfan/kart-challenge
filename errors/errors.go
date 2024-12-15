package errors

import "errors"

var (
	ErrorProductNotFound    = errors.New("product not found")
	ErrorInvalidIDSupplied  = errors.New("invalid ID supplied")
	ErrorCounonCodeNotValid = errors.New("coupon code used is not valid")
)
