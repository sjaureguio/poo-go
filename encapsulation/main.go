package encapsulation

import "fmt"

func main() {
	Sql := &Course{
		Name:    "SQL desde cero",
		Price:   159.90,
		IsFree:  false,
		UserIDs: []uint{2, 6, 8, 8},
		Classes: map[uint]string{
			1: "Select",
			2: "PL/SQL",
			3: "Store Procedures",
		},
	}

	Sql.PrintClasses()
	Sql.ChangePrice(125.60)
	fmt.Println(Sql.Price)

}
