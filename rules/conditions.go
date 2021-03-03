package rules

import "github.com/ShamanR/checkout_cart/items"

func ForSku(sku string) func(item items.ItemInterface, currentCnt int64) bool {
	return func(item items.ItemInterface, currentCnt int64) bool {
		if item.GetSKU() != sku {
			return false
		}
		return true
	}
}

func ForExactCnt(cnt int64) func(item items.ItemInterface, currentCnt int64) bool {
	return func(item items.ItemInterface, currentCnt int64) bool {
		if currentCnt != cnt {
			return false
		}
		return true
	}
}

func ForEveryCnt(cnt int64) func(item items.ItemInterface, currentCnt int64) bool {
	return func(item items.ItemInterface, currentCnt int64) bool {
		if currentCnt >= cnt {
			return true
		}
		return false
	}
}
