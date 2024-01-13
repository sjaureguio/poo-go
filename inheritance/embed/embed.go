package embed

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func NewPerson(name string, age uint8) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) Greet() {
	fmt.Println("Hello Embebbed")
}

type Student struct {
	*Person
	Grade uint8
}
