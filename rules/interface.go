package rules

import "github.com/ShamanR/checkout_cart/items"

type Rule interface {
	Discount(item items.ItemInterface, currentCnt int64) int64
	Suits(item items.ItemInterface, currentCnt int64) bool
}
