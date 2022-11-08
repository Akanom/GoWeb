package main

import (
	. "fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Println(r.Form) //print form information on server side
	Println("path", r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["url_long"])
	for l, m := range r.Form {
		Println("key:", l)
		Println("val:", strings.Join(m, ""))
	}
	Fprintf(w, "Hello Juwon!") //send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       //set router
	err := http.ListenAndServe(":9090", nil) //set listen port
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
