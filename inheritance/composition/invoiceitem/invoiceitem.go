package invoiceitem

// Item contains details of invoice
type Item struct {
	id       uint
	product  string
	quantity uint8
	value    float64
}

type Items []Item

// New returns a new Item
func New(id uint, product string, quanity uint8, value float64) Item {
	return Item{
		id:       id,
		product:  product,
		quantity: quanity,
		value:    value,
	}
}

func (i Item) Value() float64 {
	return i.value * float64(i.quantity)
}
