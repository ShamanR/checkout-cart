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
