package main

import (
	. "fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()//parse url parameters passed, then parse the response packet for post body
	Println(r.Form)//print info on the server side
	Println("path",r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["url_long"])
	for k,v:=range r.Form{
		Println("key",k)
		Println("val:",strings.Join(v,""))
	}
	Fprint(w, "Hello Juwon!")//data written to response
}

func login(w http.ResponseWriter, r *http.Request){
	Println("method:",r.Method)//get the request method
	if r.Method=="GET"{
		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		//login in the login
		Println("username",r.Form["username"])
		Println("password",r.Form["password"])
	}
}

func main(){
	http.HandleFunc("/",sayhelloName)//set router rules
	http.HandleFunc("login",login)
	err:=http.ListenAndServe(":9090",nil)//set listening port
	if err !=nil{
		log.Fatal("ListenAndServe:",err)
	}
}