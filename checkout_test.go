package checkout_cart

import (
	"github.com/ShamanR/checkout_cart/items"
	"github.com/ShamanR/checkout_cart/rules"
	"testing"
)

func TestCheckout_Total(t *testing.T) {
	tests := []struct {
		name  string
		rules []rules.Rule
		items []items.ItemInterface
		want  int64
	}{
		{
			name:  "no rules no items",
			rules: []rules.Rule{},
			items: []items.ItemInterface{},
			want:  0,
		},
		{
			name:  "no rules A item",
			rules: []rules.Rule{},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
			},
			want: 100,
		},
		{
			name:  "no rules AA items",
			rules: []rules.Rule{},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
			},
			want: 200,
		},
		{
			name:  "no rules AABBC",
			rules: []rules.Rule{},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "B", Price: 10},
				&Item{SKU: "B", Price: 10},
				&Item{SKU: "C", Price: 1},
			},
			want: 221,
		},
		{
			name: "SALE for AA=150, items AAB",
			rules: []rules.Rule{
				rules.NewRule("AA for 150").
					Condition(rules.ForSku("A"), rules.ForEveryCnt(2)).
					WillDiscount(rules.EveryCntPrice(150, 2)),
			},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "B", Price: 10},
			},
			want: 160,
		},
		{
			name: "SALE for AA=175 for AAA=150, items AAB",
			rules: []rules.Rule{
				rules.FirstOne(
					"AA=150 for AAA=200",
					rules.NewRule("AAA for 200").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(3)).
						WillDiscount(rules.EveryCntPrice(200, 3)),
					rules.NewRule("AA for 150").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(2)).
						WillDiscount(rules.EveryCntPrice(150, 2)),
				),
			},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "B", Price: 10},
			},
			want: 160,
		},
		{
			name: "SALE for AA=175 for AAA=150, items AAAB",
			rules: []rules.Rule{
				rules.FirstOne(
					"AA=150 for AAA=200",
					rules.NewRule("AAA for 200").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(3)).
						WillDiscount(rules.EveryCntPrice(200, 3)),
					rules.NewRule("AA for 150").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(2)).
						WillDiscount(rules.EveryCntPrice(150, 2)),
				),
			},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "B", Price: 10},
			},
			want: 210,
		},
		{
			name: "SALE for AA=175 for AAA=150, items AAAAB",
			rules: []rules.Rule{
				rules.FirstOne(
					"AA=150 for AAA=200",
					rules.NewRule("AAA for 200").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(3)).
						WillDiscount(rules.EveryCntPrice(200, 3)),
					rules.NewRule("AA for 150").
						Condition(rules.ForSku("A"), rules.ForEveryCnt(2)).
						WillDiscount(rules.EveryCntPrice(150, 2)),
				),
			},
			items: []items.ItemInterface{
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "A", Price: 100},
				&Item{SKU: "B", Price: 10},
			},
			want: 310,
		},
		{
			name: "readme tests",
			rules: []rules.Rule{
				rules.NewRule("Apples sale 2 for 200").
					Condition(rules.ForSku("Apples"), rules.ForEveryCnt(2)).
					WillDiscount(rules.EveryCntPrice(200, 2)),
				rules.NewRule("Bananas sale 3 for 100").
					Condition(rules.ForSku("Bananas"), rules.ForEveryCnt(3)).
					WillDiscount(rules.EveryCntPrice(100, 3)),
			},
			items: []items.ItemInterface{
				&items.Item{SKU: "Bananas", Price: 50},
				&items.Item{SKU: "Bananas", Price: 50},
				&items.Item{SKU: "Bananas", Price: 50},
				&items.Item{SKU: "Bananas", Price: 50},
				&items.Item{SKU: "Apples", Price: 150},
				&items.Item{SKU: "Apples", Price: 150},
				&items.Item{SKU: "Apples", Price: 150},
				&items.Item{SKU: "Apples", Price: 150},
			},
			want: 550,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.rules)
			for i := range tt.items {
				c.Scan(tt.items[i])
			}
			if got := c.Total(); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
