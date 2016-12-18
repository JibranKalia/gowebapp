package main

import (
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServer(":8080", nil)

}

func index(w http.Response Writer, req *http.Request){
	tpl.ExecuteTemplate
}
func about(w http.Response Writer, req *http.Request){}
func contact(w http.Response Writer, req *http.Request){}
func apply(w http.Response Writer, req *http.Request){}
