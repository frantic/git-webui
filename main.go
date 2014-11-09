package main

import (
	"fmt"
	"net/http"
	"os"
)

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	file := path + "/" + r.URL.Path[1:]
	http.ServeFile(w, r, file)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"log": ["a0372ba", "93caff1"]}`))
}

func main() {
	fmt.Println(os.Getwd())
	fmt.Println("Visit http://localhost:8080/")
	http.HandleFunc("/", frontendHandler)
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}
