package checkout_cart

import (
	"github.com/ShamanR/checkout_cart/items"
	"github.com/ShamanR/checkout_cart/rules"
)

type checkoutItem struct {
	item items.ItemInterface
	cnt  int64
}

func New(rules []rules.Rule) *Checkout {
	return &Checkout{
		rules: rules,
		items: map[string]*checkoutItem{},
	}
}

type Checkout struct {
	items map[string]*checkoutItem
	rules []rules.Rule
}

func (c *Checkout) Scan(i items.ItemInterface) {
	if _, ok := c.items[i.GetSKU()]; !ok {
		c.items[i.GetSKU()] = &checkoutItem{
			cnt:  0,
			item: i,
		}
	}
	c.items[i.GetSKU()].cnt++
}

func (c *Checkout) Total() int64 {
	total := int64(0)
	for i := range c.items {
		total += c.items[i].cnt * c.items[i].item.GetPrice()
		for r := range c.rules {
			if c.rules[r].Suits(c.items[i].item, c.items[i].cnt) {
				total += c.rules[r].Discount(c.items[i].item, c.items[i].cnt)
			}
		}
	}
	return total
}
