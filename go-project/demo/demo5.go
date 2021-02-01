package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func mydemo1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "demo 1")
}

func mydemo2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "demo 2")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", hello)

	router.HandlerFunc(http.MethodGet, "/demo1", mydemo1)
	router.Handler(http.MethodGet, "/demo2", http.HandlerFunc(mydemo2))

	log.Fatal(http.ListenAndServe(":8080", router))
}
