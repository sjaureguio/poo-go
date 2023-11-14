package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) PrintClasses() {

	c.Name = "PostgreSQL Avanzado"
	fmt.Println(c.Name)

	for _, class := range c.Classes {
		fmt.Println(class)
	}
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
