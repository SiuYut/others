package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func tmpl1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/hello.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, "template 1")
	if err != nil {
		fmt.Println(err)
	}
}

func tmpl2(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("demo.tmpl").Funcs(template.FuncMap{
		"hello": func(str string) string {
			return "Hello " + str
		},
	}).ParseFiles("./templates/demo.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, "SiuYut")
	if err != nil {
		fmt.Println(err)
	}
}

func tmpl3(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./templates/hello.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.New("hello").Parse(string(str))
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "template 3")

}

func main() {
	http.HandleFunc("/tmpl1", tmpl1)
	http.HandleFunc("/tmpl2", tmpl2)
	http.HandleFunc("/tmpl3", tmpl3)
	http.ListenAndServe(":8080", nil)
}
