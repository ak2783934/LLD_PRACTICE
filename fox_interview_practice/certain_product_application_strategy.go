package main

type CertainProductApplicationStrategy struct {
	applicableProductCategory []Item
}

func (c *CertainProductApplicationStrategy) isApplicable(order Order, user User) bool {
	for _, item := range order.items {
		if !SearchInArray(c.applicableProductCategory, item.category) {
			return false
		}
	}
	return true
}

func SearchInArray(array []string, key string) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}
	return false
}
