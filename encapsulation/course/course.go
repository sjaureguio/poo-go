package course

import "fmt"

type Course struct {
	name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	classes map[uint]string
}

func New(name string, price float64, isFree bool) *Course {
	if price == 0 {
		price = 30.0
	}

	return &Course{
		name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *Course) SetName(name string) {
	c.name = name
}

func (c *Course) Name() string {
	return c.name
}

func (c *Course) SetClasses(classes map[uint]string) {
	c.classes = classes

}

func (c *Course) PrintClasses() {
	for _, class := range c.classes {
		fmt.Println(class)
	}
}
