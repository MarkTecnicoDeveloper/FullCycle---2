package main

import (
	"html/template"
    "net/http"
)

func main() {
   http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil) 
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"CodeEducation": greeting("Code.education Rocks!!!"),
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}
func greeting(str string)string{
	return str 
}