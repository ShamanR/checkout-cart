package rules

import "github.com/ShamanR/checkout_cart/items"

// Returns new price for item
// Similar to newPricePerItem*cnt
func Price(newPricePerItem int64) func(item items.ItemInterface, currentCnt int64) int64 {
	return func(item items.ItemInterface, currentCnt int64) int64 {
		return newPricePerItem*currentCnt - currentCnt*item.GetPrice()
	}
}

// Returns Discount Calculator
// every CNT items in the cart will have newPrice, all others will have ordinary price
func EveryCntPrice(newPrice int64, cnt int64) func(item items.ItemInterface, currentCnt int64) int64 {
	if cnt == 0 {
		panic("impossible to create rule with cnt 0")
	}
	return func(item items.ItemInterface, currentCnt int64) int64 {
		discountCnt := currentCnt / cnt
		ordinaryCnt := currentCnt - discountCnt*cnt
		return discountCnt*newPrice + ordinaryCnt*item.GetPrice() - currentCnt*item.GetPrice()
	}
}
