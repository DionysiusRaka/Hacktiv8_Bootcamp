package database

import (
	"fmt"

	_ "github.com/lib/pq"
)

type Employees struct {
	ID       string
	FullName string
	Email    string
	Age      int
	Division string
}

func CreateEmployees() {
	var employees = Employees{}

	sql := `
		INSERT INTO employees (full_name,email,age,division)
		VALUES ($1,$2,$3,$4)
		RETURNING *
	`

	err = db.QueryRow(sql, "Isnanda", "afin@gmail.com", 25, "IT").Scan(&employees.ID, &employees.FullName, &employees.Email, &employees.Age, &employees.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employees)
}

func GetEmployees() {
	var employees = []Employees{}

	sql := `
	SELECT * FROM employees
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employees{}

		err = rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}

	fmt.Println("Employee Data :", employees)
}
