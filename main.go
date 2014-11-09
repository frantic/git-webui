package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	ui, _ := ioutil.ReadFile("log.html")
	w.Write(ui)
}

func main() {
	fmt.Println("Visit http://localhost:8080/log")
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}
