package main

import (
	. "fmt"
	"net/http"
)

type MyMux struct{
}

func (p *MyMux) ServerHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path=="/"{
		sayhelloName(w, r)
		return
	}
	http.NotFound(w,r)
	return
}


func sayhelloName(w http.ResponseWriter, r *http.Request){
	Fprintf(w,"Hello Juwon")
}

func main(){
	mux:=&MyMux{}
	http.ListenAndServe(":9090",mux)
}