package main

import (
	"html/template"
	"net/http"
)

var users []*User

type User struct {
	Username string
	Password string
}

// info ro az karbar migirim
func login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// info ro be string users ezafe mikonim
func add(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user")
	password := r.FormValue("pass")
	text := User{Username: username, Password: password}
	users = append(users, &text)
	http.Redirect(w, r, "/list/", http.StatusFound)
}

// info ro be karbar neshoon midim
func list(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "list.html", users)
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/add/", add)
	http.HandleFunc("/list/", list)
	// http.HandleFunc("/", login)
	http.ListenAndServe(":80", nil)
}
