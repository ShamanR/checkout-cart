package items

type ItemInterface interface {
	GetSKU() string
	GetPrice() int64
}

type Item struct {
	SKU   string
	Price int64
}

func (i *Item) GetSKU() string {
	return i.SKU
}

func (i *Item) GetPrice() int64 {
	return i.Price
}
