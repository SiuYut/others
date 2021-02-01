package main

import (
	"fmt"
	"net/http"
)

type DemoHandler struct {}

func (h *DemoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Golang!")
}

func main() {
	
	//handler := DemoHandler{}
	//server := http.Server{
	//	Addr: ":8080",
	//	Handler: &handler,
	//}

	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/demo/", &DemoHandler{})

	server.ListenAndServe()
}
