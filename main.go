package main

import (
	"fmt"
	"github.com/sjaureguio/poo-go/encapsulation/course"
	"github.com/sjaureguio/poo-go/inheritance/composition/customer"
	"github.com/sjaureguio/poo-go/inheritance/composition/invoice"
	"github.com/sjaureguio/poo-go/inheritance/composition/invoiceitem"
	"github.com/sjaureguio/poo-go/inheritance/embed"
	"github.com/sjaureguio/poo-go/inheritance/override"
)

func main() {
	Go := course.New("POO en Go", 0, false)
	Go.UserIDs = []uint{2, 6, 8, 8}
	Go.SetClasses(map[uint]string{
		1: "Select",
		2: "PL/SQL",
		3: "Store Procedures",
	})
	Go.PrintClasses()

	// Composición
	cust := customer.New("Solano Jauregui", "Calle REAL 225", "+51939496226")
	items := []invoiceitem.Item{
		invoiceitem.New(1, "Coca cola 1.5LT", 1, 50),
		invoiceitem.New(2, "Galleta Ritz", 1, 5),
	}
	i := invoice.New(1, "2023-11-14", cust, items)
	i.SetTotal()

	fmt.Printf("%+v \n", i)

	// Embeber Struct
	s := embed.Student{
		Person: embed.NewPerson("Solano Jauregui", 25),
		Grade:  9,
	}
	s.Greet()
	fmt.Println(s.Name)
	fmt.Println(s.Person.Name)

	// Embeber interface
	c := override.CompoundForm{
		Form: override.Square{Side: 5},
	}

	area := c.Area()
	fmt.Println(area)
}
