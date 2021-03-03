package rules

import "github.com/ShamanR/checkout_cart/items"

func Price(newPrice int64) func(item items.ItemInterface, currentCnt int64) int64 {
	return func(item items.ItemInterface, currentCnt int64) int64 {
		return newPrice*currentCnt - currentCnt*item.GetPrice()
	}
}

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
