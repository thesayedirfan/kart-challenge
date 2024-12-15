package utils

import (
	"strconv"
	"strings"

	"github.com/thesayedirfan/kart-challenge/errors"
)

func GetDiscountedPrice(totalamount float64, percentage float64) float64 {
	return totalamount * (percentage / 100)
}

func ValidateDiscountCoupon(coupon string, coupons map[string]int) (bool, error) {
	count, exist := coupons[coupon]
	if count >= 2 && exist {
		return true, nil
	}
	return false, errors.ErrorCounonCodeNotValid
}

func IsValidProductID(s string) bool {
	s = strings.TrimSpace(s)
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}
