package main

import (
	"fmt"
	"html/template" //template html
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	//define template
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// define route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	//print tem
	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)
}
