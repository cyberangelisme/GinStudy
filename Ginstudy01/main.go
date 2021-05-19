package main

import (
	"fmt"
	"net/http"
	"text/template"
	//"github.com/gin-gonic/gin"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "小娃子"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "小娃子"
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html", "./templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "小娃子"
	t.ExecuteTemplate(w, "index.html", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html", "./templates/home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "小娃子"
	t.ExecuteTemplate(w, "home.html", msg)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	http.ListenAndServe(":9090", nil)

}
