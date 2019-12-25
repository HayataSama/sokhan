package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct{}

var users []*user

type user struct {
	username string
	password string
}

// info ro az karbar migirim
func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	p := &Page{}
	t.Execute(w, p)
}

// info ro be string users ezafe mikonim
func add(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user")
	password := r.FormValue("pass")
	text := &user{username: username, password: password}
	users = append(users, text)
	http.Redirect(w, r, "/list/", http.StatusFound)
}

// info ro be karbar neshoon midim
func list(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("list.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	p := &Page{}
	t.Execute(w, p)

	fmt.Fprintf(w, "<ul>")
	for i := 0; i < len(users); i++ {
		fmt.Fprintf(w, "<li>"+"%s"+"<br>"+"%s"+"</li><br>", users[i].username, users[i].password)
	}

	fmt.Fprintf(w, "</ul><br><a href=\"/\">Back to main page</a>")
}

func main() {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/add/", add)
	http.HandleFunc("/list/", list)
	// http.HandleFunc("/", login)
	http.ListenAndServe(":80", nil)
}
