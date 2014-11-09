package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	ui, _ := ioutil.ReadFile("log.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write(ui)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"log": ["a0372ba", "93caff1"]}`))
}

func main() {
	fmt.Println("Visit http://localhost:8080/")
	http.HandleFunc("/", frontendHandler)
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}
