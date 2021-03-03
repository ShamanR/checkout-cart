# Simple Checkout calculator 
[![GoDoc](https://godoc.org/github.com/ShamanR/checkout-cart?status.png)](https://godoc.org/github.com/github.com/ShamanR/checkout-cart)
[![Go](https://github.com/ShamanR/checkout_cart/actions/workflows/go.yml/badge.svg)](https://github.com/ShamanR/checkout_cart/actions/workflows/go.yml)

It is a simple test project It allows you to create market checkout cart And add different kinds of adv campaigns

## Simple example

```golang

bananasItem := items.Item{SKU:"Bananas", Price:50}
apples := items.Item{SKU:"Apples", Price:150}

ch := New([]simple_checkout.Rule{
rules.NewRule("Apples sale 2 for 200").
Condition(rules.ForSku("Apples"), rules.ForEveryCnt(2)).
WillDiscount(rules.EveryCntPrice(200, 2),
),
rules.NewRule("Bananas sale 3 for 100").
Condition(rules.ForSku("Bananas"), rules.ForEveryCnt(3)).
WillDiscount(rules.EveryCntPrice(300, 3),
),
})
// scan 3 bananas: 3 with sale for 100$, last for 50
ch.Scan(bananasItem)
ch.Scan(bananasItem)
ch.Scan(bananasItem)
ch.Scan(bananasItem)
// scan 4 apples: 2 sale pair for 200 each
ch.Scan(apples)
ch.Scan(apples)
ch.Scan(apples)
ch.Scan(apples)

total := ch.Total() // 550
```

## Install

