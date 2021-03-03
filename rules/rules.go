package rules

import "github.com/ShamanR/checkout_cart/items"

// Creates rule
func NewRule(name string) *rule {
	s := func(item items.ItemInterface, currentCnt int64) bool { return false }
	d := func(item items.ItemInterface, currentCnt int64) int64 { return int64(0) }
	return &rule{
		Name:     name,
		suits:    []func(item items.ItemInterface, currentCnt int64) bool{s},
		discount: d,
	}
}

type rule struct {
	Name     string
	suits    []func(item items.ItemInterface, currentCnt int64) bool
	discount func(item items.ItemInterface, currentCnt int64) int64
}

// Add condition to the rule
func (r *rule) Condition(suits ...func(item items.ItemInterface, currentCnt int64) bool) *rule {
	r.suits = suits
	return r
}

// Set Discount to the rule
func (r *rule) WillDiscount(discount func(item items.ItemInterface, currentCnt int64) int64) *rule {
	r.discount = discount
	return r
}

// Check, if item suits all conditions
func (r *rule) Suits(item items.ItemInterface, currentCnt int64) bool {
	for i := range r.suits {
		if !r.suits[i](item, currentCnt) {
			return false
		}
	}
	return true
}

// Calc item discount
func (r *rule) Discount(item items.ItemInterface, currentCnt int64) int64 {
	return r.discount(item, currentCnt)
}

// Create rule based on other rules
// The first rule with condition==true will be used
// Similar to logic OR
func FirstOne(name string, rules ...Rule) *rule {
	r := &rule{
		Name: name,
	}

	cond := func(item items.ItemInterface, currentCnt int64) bool {
		for i := range rules {
			if rules[i].Suits(item, currentCnt) {
				r.WillDiscount(rules[i].Discount)
				return true
			}
		}
		return false
	}
	r.Condition(cond)
	return r
}
