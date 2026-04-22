package main

import "sync"

type CouponStatus string

const (
	ACTIVE    CouponStatus = "ACTIVE"
	EXPIRED   CouponStatus = "EXPIRED"
	IN_ACTIVE CouponStatus = "IN_ACTIVE"
)

type Coupon struct {
	id                  int
	usage               int
	maxUsage            int
	discountStrategy    DiscountStrategy
	applicationStrategy ApplicationStrategy
	status              CouponStatus
	mu                  sync.Mutex
}
