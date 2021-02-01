package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "demo 1")
}

func demo2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "demo 2")
}

func demo3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "demo 3")
}

func main() {
	http.Handle("/demo1", &MyHandler{})
	http.Handle("/demo2", http.HandlerFunc(demo2))
	http.HandleFunc("/demo3", demo3)
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe("", nil)
}
