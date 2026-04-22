Design a Coupon & Discount Engine for an e-commerce platform. 
The system should allow the operations team to create, manage, and apply discount coupons to customer orders.

> Core Requirements:
> 1. Support multiple discount types: eg: flat amount off, percentage off, and buy X get Y free
> 2. Each coupon can have applicability rules — e.g., applicable only to certain products, categories, or user segments
> 3. Each coupon can have constraints — max usage count, per-user usage limit, minimum cart value, valid date range
> 4. A user should be able to apply a coupon to their cart and see the discounted total 
> 5. The system should be extensible — adding a new discount type or a new applicability rule should require minimal changes



Entities: 

Coupon ---
CouponManager  ------
DiscountStrategy ---
ApplicationStrategy ---
User ---
Order ---