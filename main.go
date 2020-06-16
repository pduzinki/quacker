package main

import (
	"net/http"

	"bytes"
	"html/template"
	"io"

	"github.com/gorilla/mux"
)

// Homepage handles quacker homepage
func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("views/layouts/page.gohtml", "views/layouts/navbar.gohtml")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.ExecuteTemplate(&buf, "page", nil)
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

// Login handles GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// LoginPost handles POST /login
func LoginPost(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// Signup handles GET /signup
func Signup() {
	// TODO
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Homepage)

	http.ListenAndServe(":3000", r)
}
