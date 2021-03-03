package rules

import "github.com/ShamanR/checkout_cart/items"

// describe pricing rules
type Rule interface {
	// If current item with current CNT suits the rule
	Suits(item items.ItemInterface, currentCnt int64) bool
	// If Item suits (see Suits method) Discount method calculates discount
	// Discount will be add to ordinary cost(price*cnt)
	// That's why Discount should return negative for Sale discounts
	Discount(item items.ItemInterface, currentCnt int64) int64
}
