package invoice

import (
	"github.com/sjaureguio/poo-go/inheritance/composition/customer"
	"github.com/sjaureguio/poo-go/inheritance/composition/invoiceitem"
)

// Invoice is a struct of Invoice
type Invoice struct {
	id          uint
	dateInvoice string
	total       float64
	customer    customer.Customer
	items       []invoiceitem.Item
}

// New return a new Invoice
func New(id uint, dateInvoice string, cust customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		id:          id,
		dateInvoice: dateInvoice,
		customer:    cust,
		items:       items,
	}
}

// SetTotal is a setter of invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
