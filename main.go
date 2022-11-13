package main

import (
	"crypto/md5"
	. "fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()   //parse url parameters passed, then parse the response packet for post body
	Println(r.Form) //print info on the server side
	Println("path", r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["url_long"])
	for k, v := range r.Form {
		Println("key", k)
		Println("val:", strings.Join(v, ""))
	}
	Fprint(w, "Hello Juwon!") //data written to response
}

func login(w http.ResponseWriter, r *http.Request) {
	Println("method:", r.Method) //get the request method
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		//log in request
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//check token availiability
		} else {
			//return error if the token is not availiable
		}
		//login in the login
		Println("username length:", len(r.Form["username"][0]))
		Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //HTMLEscapestring allows cross scripting
		Println("password", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //set router rules
	http.HandleFunc("login", login)
	err := http.ListenAndServe(":9090", nil) //set listening port
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
