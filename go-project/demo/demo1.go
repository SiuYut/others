package main

import (
	"fmt"
	"net/http"
)

func DemoHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s  %s  %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
	//fmt.Fprintf(w, "%s\n", r.Header["Accept-Language"])
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
}

func DemoHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s  %s  %s\n", r.Method, r.URL, r.Proto)
	for k, _ := range r.URL.Query() {
		fmt.Fprintf(w, "parameter %q, value %q\n", k, r.URL.Query().Get(k))
	}
}

func DemoHandler3(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	gender := r.FormValue("gender")
	fmt.Fprintf(w, "Username: %s  Gender: %s\n", username, gender)
}

func DemoHandler4(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "Field: %s  Value: %s\n", key, value)
		//fmt.Fprintf(w, "Field: %s  Value: %s\n", key, r.FormValue(key))
	}
}

func DemoHandler5(w http.ResponseWriter, r *http.Request) {
	for _,cookie := range r.Cookies() {
		fmt.Fprintf(w, "cookie: %s  value: %s\n", cookie.Name, cookie.Value)
	}
}

func main() {
	http.HandleFunc("/demo", DemoHandler1)
	http.ListenAndServe(":8080", nil)
}
