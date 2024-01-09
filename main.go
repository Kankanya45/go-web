package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rt := mux.NewRouter()

	//ex: http://localhost:8080/books/cprogramming/page/1
	//title = cprogramming
	// page = 1

	rt.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// ex : http://localhost:80
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", rt)
}
