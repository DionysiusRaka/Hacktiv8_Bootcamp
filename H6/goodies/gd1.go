package goodies

import (
	"fmt"
	"reflect"
)

type Organizations struct {
	orgName string
	address string
}

type Employee struct {
	nik      string
	name     string
	age      int
	division string
	address  Organizations
}

func Ref2Struct() {
	employees := []Employee{
		{
			nik:      "12398765",
			name:     "<NAME>",
			age:      34,
			division: "IT",
			address: Organizations{
				orgName: "PT. Makmur",
				address: "Jl. Makmur No. 1",
			},
		},
		{
			nik:      "1234567",
			name:     "<NAME>",
			age:      33,
			division: "IT WarWor",
			address: Organizations{
				orgName: "PT. Makmur 2",
				address: "Jl. Makmur No. 5",
			},
		},
		{
			nik:      "1234567",
			name:     "<NAME>",
			age:      38,
			division: "IT WarWor",
			address: Organizations{
				orgName: "PT. Makmur 2",
				address: "Jl. Makmur No. 5",
			},
		},
	}
	for _, employee := range employees {
		fmt.Println(reflect.ValueOf(employees), reflect.TypeOf(employee))
	}

	fmt.Println(findOldest(employees))

}
func findOldest(employees []Employee) Employee {
	oldestAge := -1
	var oldestEmployee Employee
	for _, employee := range employees {
		if employee.age > oldestAge {
			oldestAge = employee.age
			oldestEmployee = employee
		}
	}
	return oldestEmployee
}
