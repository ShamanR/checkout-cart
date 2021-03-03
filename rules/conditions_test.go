package rules

import (
	"github.com/ShamanR/checkout_cart/items"
	"reflect"
	"testing"
)

func TestForSku(t *testing.T) {
	tests := []struct {
		name string
		sku  string
		item items.ItemInterface
		want bool
	}{
		{
			name: "empty",
			sku:  "",
			item: &items.Item{SKU: "sku"},
			want: false,
		},
		{
			name: "sku equals",
			sku:  "sku",
			item: &items.Item{SKU: "sku"},
			want: true,
		},
		{
			name: "sku differs",
			sku:  "sku1",
			item: &items.Item{SKU: "sku2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cond := ForSku(tt.sku)
			if got := cond(tt.item, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForSku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForExactCnt(t *testing.T) {
	tests := []struct {
		name           string
		exactCnt       int64
		item           items.ItemInterface
		currentItemCnt int64
		want           bool
	}{
		{
			name:           "zero tolerance",
			exactCnt:       0,
			item:           &items.Item{},
			currentItemCnt: 0,
			want:           true,
		},
		{
			name:           "Cnt equals",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 4,
			want:           true,
		},
		{
			name:           "currentItemCnt smaller",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 1,
			want:           false,
		},
		{
			name:           "currentItemCnt bigger",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 10,
			want:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cond := ForExactCnt(tt.exactCnt)
			if got := cond(tt.item, tt.currentItemCnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForExactCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEveryCnt(t *testing.T) {
	tests := []struct {
		name           string
		exactCnt       int64
		item           items.ItemInterface
		currentItemCnt int64
		want           bool
	}{
		{
			name:           "zero tolerance",
			exactCnt:       0,
			item:           &items.Item{},
			currentItemCnt: 0,
			want:           true,
		},
		{
			name:           "Cnt equals",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 4,
			want:           true,
		},
		{
			name:           "currentItemCnt smaller",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 1,
			want:           false,
		},
		{
			name:           "currentItemCnt bigger",
			exactCnt:       4,
			item:           &items.Item{},
			currentItemCnt: 10,
			want:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cond := ForEveryCnt(tt.exactCnt)
			if got := cond(tt.item, tt.currentItemCnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForEveryCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}
