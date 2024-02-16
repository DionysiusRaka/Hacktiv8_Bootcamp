package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)
	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("Server started at localhost:", server.Addr)
	server.ListenAndServe()

}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudent())
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET requests are allowed"))
		return false
	}
	return true
}
