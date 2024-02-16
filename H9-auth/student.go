package main

import (
	"encoding/json"
	"net/http"
)

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetStudent() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{
		Id:    "1",
		Name:  "Nanda Lemon",
		Grade: 2,
	})
	students = append(students, &Student{
		Id:    "2",
		Name:  "Azwar Reno",
		Grade: 2,
	})
	students = append(students, &Student{
		Id:    "3",
		Name:  "Afin Kopling",
		Grade: 3,
	})
}
