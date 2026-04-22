package main

type CouponManager struct {
	Coupons map[int]*Coupon
	count   int
}

func CreateCouponManager() *CouponManager {
	return &CouponManager{
		Coupons: make(map[int]*Coupon, 0),
		count:   0,
	}
}

func (c *CouponManager) AddCoupons(discountStrategy DiscountStrategy, applicationStrategy ApplicationStrategy, maxUsage int) {
	coupon := &Coupon{
		id:                  c.count,
		usage:               0,
		maxUsage:            maxUsage,
		status:              ACTIVE,
		discountStrategy:    discountStrategy,
		applicationStrategy: applicationStrategy,
	}
	c.Coupons[c.count] = coupon
	c.count++
}

func (c *CouponManager) UpdateCouponStatus(id int, newStatus CouponStatus) {
	coupon, ok := c.Coupons[id]
	if !ok {
		return
	}
	coupon.status = newStatus
}

func (c *CouponManager) UpdateCouponCount(id int, newCount int) {
	if newCount >= 0 {
		coupon, ok := c.Coupons[id]
		if !ok {
			return
		}
		coupon.usage = newCount
	}
}
