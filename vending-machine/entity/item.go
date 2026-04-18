package entity

type Item struct {
	name  string
	code  string
	price int
}

func NewItem(name, code string, price int) *Item {
	return &Item{
		name:  name,
		code:  code,
		price: price,
	}
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetCode() string {
	return i.code
}

func (i *Item) GetPrice() int {
	return i.price
}
