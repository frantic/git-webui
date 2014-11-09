package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Visit http://localhost:8080/git")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
