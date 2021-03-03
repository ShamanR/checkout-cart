package rules

import (
	"github.com/ShamanR/checkout_cart/items"
	"reflect"
	"testing"
)

func TestPrice(t *testing.T) {
	type args struct {
		item items.ItemInterface
		cnt  int64
	}
	tests := []struct {
		name     string
		newPrice int64
		args     args
		discount int64
	}{
		{
			name:     "cnt=0, oldPrice=200, newPrice=100",
			newPrice: 100,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  0,
			},
			discount: 0,
		},
		{
			name:     "cnt=1, oldPrice=200, newPrice=100",
			newPrice: 100,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  1,
			},
			discount: -100,
		},
		{
			name:     "cnt=1, oldPrice=200, newPrice=10",
			newPrice: 10,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  1,
			},
			discount: -190,
		},
		{
			name:     "cnt=2, oldPrice=200, newPrice=10",
			newPrice: 10,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  2,
			},
			discount: -380,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price := Price(tt.newPrice)
			if got := price(tt.args.item, tt.args.cnt); !reflect.DeepEqual(got, tt.discount) {
				t.Errorf("Price() = %v, discount %v", got, tt.discount)
			}
		})
	}
}

func TestEveryCntPrice(t *testing.T) {
	type args struct {
		item items.ItemInterface
		cnt  int64
	}
	tests := []struct {
		name     string
		newPrice int64
		everyCnt int64
		args     args
		discount int64
	}{
		{
			name:     "for each 1 item get price 100 instead of 200",
			newPrice: 100,
			everyCnt: 1,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  1,
			},
			discount: -100,
		},
		{
			name:     "cnt = 0",
			newPrice: 100,
			everyCnt: 1,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  0,
			},
			discount: 0,
		},
		{
			name:     "for each 1 item get price 100 instead of 200",
			newPrice: 100,
			everyCnt: 1,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  5,
			},
			discount: -500,
		},
		{
			name:     "for each 1 item get price 100 instead of 200",
			newPrice: 100,
			everyCnt: 2,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  3,
			},
			discount: -300,
		},
		{
			name:     "for each 1 item get price 100 instead of 200",
			newPrice: 100,
			everyCnt: 2,
			args: args{
				item: &items.Item{Price: 200},
				cnt:  4,
			},
			discount: -600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discountRule := EveryCntPrice(tt.newPrice, tt.everyCnt)
			if got := discountRule(tt.args.item, tt.args.cnt); !reflect.DeepEqual(got, tt.discount) {
				t.Errorf("discount: got() = %v, want %v", got, tt.discount)
			}
		})
	}
}
