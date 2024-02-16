package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var PORT = ":8080"

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

func MainApi() {
	http.HandleFunc("/employee", GetEmployees)
	http.HandleFunc("/employees", CreateEmployees)
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

var Employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 23, Division: "Finance"}, {ID: 3, Name: "Mailo", Age: 20, Division: "IT"},
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, Employees)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		divison := r.FormValue("division")
		convertAge, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		newEmployee := Employee{
			ID:       len(Employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: divison,
		}
		Employees = append(Employees, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}
